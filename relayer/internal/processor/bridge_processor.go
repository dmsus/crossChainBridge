package processor

import (
    "context"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "math/big"

    "github.com/dmsus/crossChainBridge/relayer/internal/eventlistener"
    "github.com/dmsus/crossChainBridge/relayer/internal/sender"
    "github.com/dmsus/crossChainBridge/relayer/pkg/database"
    "github.com/ethereum/go-ethereum/common"
)

// BridgeProcessor обрабатывает события с гарантией идемпотентности
type BridgeProcessor struct {
    polygonSender *sender.PolygonSender
    repository    *database.Repository
    maxRetries    int
}

// Config конфигурация для процессора
type Config struct {
    PolygonSender *sender.PolygonSender
    Repository    *database.Repository
    MaxRetries    int
}

// NewBridgeProcessor создает новый процессор
func NewBridgeProcessor(cfg Config) *BridgeProcessor {
    return &BridgeProcessor{
        polygonSender: cfg.PolygonSender,
        repository:    cfg.Repository,
        maxRetries:    cfg.MaxRetries,
    }
}

// ProcessEvent обрабатывает событие с гарантией идемпотентности
func (p *BridgeProcessor) ProcessEvent(ctx context.Context, event *eventlistener.TokensLockedEvent) error {
    // Генерируем уникальный хеш события для идемпотентности
    eventHash := p.generateEventHash(event)
    
    log.Printf("🔍 Processing event with idempotency: hash=%s", eventHash)

    // 1. Проверяем идемпотентность
    isNew, err := p.repository.CheckAndStoreEvent(ctx, eventHash, int64(event.BlockNumber), 0, event.User.Hex())
    if err != nil {
        return fmt.Errorf("failed to check event idempotency: %v", err)
    }
    
    if !isNew {
        log.Printf("⏭️ Event already processed, skipping: %s", eventHash)
        return nil
    }

    // 2. Проверяем nonce для защиты от replay атак
    nonceValid, err := p.repository.CheckNonce(ctx, event.User.Hex(), event.TargetChainID.Int64(), event.Nonce.Int64())
    if err != nil {
        return fmt.Errorf("failed to check nonce: %v", err)
    }
    
    if !nonceValid {
        return fmt.Errorf("invalid nonce, possible replay attack: user=%s, nonce=%s", 
            event.User.Hex(), event.Nonce.String())
    }

    // 3. Создаем запись о транзакции
    transaction, err := p.repository.CreateTransaction(ctx,
        event.Nonce.Int64(),
        event.User.Hex(),
        event.Amount.String(),
        11155111, // Ethereum Sepolia chain ID
        event.TargetChainID.Int64(), // target chain (Polygon)
        event.TargetAddress.Hex(),
        "", // signature будет сгенерирована позже
    )
    if err != nil {
        return fmt.Errorf("failed to create transaction: %v", err)
    }

    // 4. Обновляем статус на processing
    if err := p.repository.UpdateTransactionStatus(ctx, transaction.ID, database.StateProcessing, "", ""); err != nil {
        return fmt.Errorf("failed to update transaction status: %v", err)
    }

    // 5. Обрабатываем транзакцию
    return p.processTransaction(ctx, transaction, event)
}

// processTransaction обрабатывает отдельную транзакцию
func (p *BridgeProcessor) processTransaction(ctx context.Context, transaction *database.Transaction, event *eventlistener.TokensLockedEvent) error {
    var lastError error
    
    for retry := 0; retry <= p.maxRetries; retry++ {
        if retry > 0 {
            log.Printf("🔄 Retry attempt %d/%d for transaction %d", retry, p.maxRetries, transaction.ID)
        }

        // Отправляем транзакцию в Polygon
        userAddr := common.HexToAddress(transaction.UserAddress)
        amount := new(big.Int)
        amount.SetString(transaction.Amount, 10)
        nonce := big.NewInt(transaction.EventNonce)

        tx, err := p.polygonSender.SendReleaseTokens(ctx, userAddr, amount, nonce)
        if err != nil {
            lastError = err
            log.Printf("❌ Transaction failed (attempt %d): %v", retry+1, err)
            
            // Если это последняя попытка, обновляем статус на failed
            if retry == p.maxRetries {
                updateErr := p.repository.UpdateTransactionStatus(ctx, transaction.ID, database.StateFailed, "", err.Error())
                if updateErr != nil {
                    log.Printf("⚠️ Failed to update transaction status: %v", updateErr)
                }
                return fmt.Errorf("transaction failed after %d retries: %v", p.maxRetries, err)
            }
            
            continue
        }

        // Транзакция успешно отправлена
        if err := p.repository.UpdateTransactionStatus(ctx, transaction.ID, database.StateCompleted, tx.Hash().Hex(), ""); err != nil {
            return fmt.Errorf("failed to update transaction status to completed: %v", err)
        }

        log.Printf("🎉 Transaction completed successfully: id=%d, txHash=%s", transaction.ID, tx.Hash().Hex())
        return nil
    }

    return lastError
}

// generateEventHash генерирует уникальный хеш события для идемпотентности
func (p *BridgeProcessor) generateEventHash(event *eventlistener.TokensLockedEvent) string {
    data := fmt.Sprintf("%s:%s:%s:%s:%d",
        event.User.Hex(),
        event.Amount.String(),
        event.Nonce.String(),
        event.TargetAddress.Hex(),
        event.BlockNumber)
    
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

// RecoverPendingTransactions восстанавливает pending транзакции при запуске
func (p *BridgeProcessor) RecoverPendingTransactions(ctx context.Context) error {
    log.Println("🔍 Recovering pending transactions...")
    
    pendingTransactions, err := p.repository.GetPendingTransactions(ctx)
    if err != nil {
        return fmt.Errorf("failed to get pending transactions: %v", err)
    }
    
    if len(pendingTransactions) == 0 {
        log.Println("✅ No pending transactions to recover")
        return nil
    }
    
    log.Printf("🔄 Found %d pending transactions to recover", len(pendingTransactions))
    
    for _, transaction := range pendingTransactions {
        log.Printf("🔄 Recovering transaction: id=%d, nonce=%d, user=%s", 
            transaction.ID, transaction.EventNonce, transaction.UserAddress)
        
        // Создаем fake event для восстановления
        event := &eventlistener.TokensLockedEvent{
            User:          common.HexToAddress(transaction.UserAddress),
            Amount:        new(big.Int),
            Nonce:         big.NewInt(transaction.EventNonce),
            TargetChainID: big.NewInt(transaction.TargetChainID),
            TargetAddress: common.HexToAddress(transaction.TargetAddress),
            BlockNumber:   uint64(transaction.SourceChainID), // временно используем source chain как block number
        }
        event.Amount.SetString(transaction.Amount, 10)
        
        // Обрабатываем транзакцию заново
        if err := p.processTransaction(ctx, &transaction, event); err != nil {
            log.Printf("❌ Failed to recover transaction %d: %v", transaction.ID, err)
        }
    }
    
    return nil
}

// HealthCheck проверяет здоровье процессора
func (p *BridgeProcessor) HealthCheck(ctx context.Context) error {
    return p.repository.HealthCheck(ctx)
}
