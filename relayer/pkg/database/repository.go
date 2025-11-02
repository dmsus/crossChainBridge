package database

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/lib/pq"
)

// TransactionState представляет состояние транзакции
type TransactionState string

const (
    StatePending    TransactionState = "pending"
    StateProcessing TransactionState = "processing"
    StateCompleted  TransactionState = "completed"
    StateFailed     TransactionState = "failed"
    StateRetry      TransactionState = "retry"
)

// Transaction представляет кросс-чейн транзакцию
type Transaction struct {
    ID              int64
    EventNonce      int64
    UserAddress     string
    Amount          string
    SourceChainID   int64
    TargetChainID   int64
    TargetAddress   string
    Status          TransactionState
    Signature       sql.NullString
    CreatedAt       time.Time
    UpdatedAt       time.Time
}

// ProcessedEvent представляет обработанное событие для идемпотентности
type ProcessedEvent struct {
    ID              int64
    EventHash       string
    BlockNumber     int64
    LogIndex        int64
    ContractAddress string
    TransactionID   sql.NullInt64
    ProcessedAt     time.Time
}

// NonceTracking отслеживает nonce для защиты от replay атак
type NonceTracking struct {
    ID          int64
    UserAddress string
    ChainID     int64
    LastNonce   int64
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Repository предоставляет методы для работы с БД
type Repository struct {
    db *sql.DB
}

// NewRepository создает новый репозиторий
func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

// CheckAndStoreEvent проверяет идемпотентность и сохраняет событие если оно новое
func (r *Repository) CheckAndStoreEvent(ctx context.Context, eventHash string, blockNumber, logIndex int64, contractAddress string) (bool, error) {
    var exists bool
    err := r.db.QueryRowContext(ctx,
        `SELECT EXISTS(SELECT 1 FROM processed_events WHERE event_hash = $1)`,
        eventHash).Scan(&exists)
    
    if err != nil {
        return false, fmt.Errorf("failed to check event existence: %v", err)
    }
    
    if exists {
        log.Printf("⚠️ Event already processed: %s", eventHash)
        return false, nil
    }
    
    // Вставляем запись о обработанном событии
    _, err = r.db.ExecContext(ctx,
        `INSERT INTO processed_events (event_hash, block_number, log_index, contract_address, processed_at)
         VALUES ($1, $2, $3, $4, NOW())`,
        eventHash, blockNumber, logIndex, contractAddress)
    
    if err != nil {
        // Если это нарушение уникальности по (block_number, log_index, contract_address)
        if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
            log.Printf("⚠️ Event already processed (unique constraint): %s", eventHash)
            return false, nil
        }
        return false, fmt.Errorf("failed to insert processed event: %v", err)
    }
    
    log.Printf("✅ Event stored for idempotency: %s", eventHash)
    return true, nil
}

// CreateTransaction создает новую транзакцию
func (r *Repository) CreateTransaction(ctx context.Context, eventNonce int64, userAddress string, amount string, sourceChainID, targetChainID int64, targetAddress, signature string) (*Transaction, error) {
    var transaction Transaction
    
    err := r.db.QueryRowContext(ctx,
        `INSERT INTO transactions (event_nonce, user_address, amount, source_chain_id, target_chain_id, target_address, status, signature, created_at, updated_at)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
         RETURNING id, event_nonce, user_address, amount, source_chain_id, target_chain_id, target_address, status, signature, created_at, updated_at`,
        eventNonce, userAddress, amount, sourceChainID, targetChainID, targetAddress, StatePending, signature).Scan(
        &transaction.ID, &transaction.EventNonce, &transaction.UserAddress, &transaction.Amount,
        &transaction.SourceChainID, &transaction.TargetChainID, &transaction.TargetAddress, 
        &transaction.Status, &transaction.Signature, &transaction.CreatedAt, &transaction.UpdatedAt)
    
    if err != nil {
        return nil, fmt.Errorf("failed to create transaction: %v", err)
    }
    
    log.Printf("✅ Transaction created: id=%d, nonce=%d, user=%s", 
        transaction.ID, transaction.EventNonce, transaction.UserAddress)
    
    return &transaction, nil
}

// UpdateTransactionStatus обновляет статус транзакции
func (r *Repository) UpdateTransactionStatus(ctx context.Context, transactionID int64, status TransactionState, txHash string, errorMsg string) error {
    // В существующей схеме нет полей tx_hash и error_message, обновляем только статус
    result, err := r.db.ExecContext(ctx,
        `UPDATE transactions SET status = $1, updated_at = NOW() WHERE id = $2`,
        status, transactionID)
    
    if err != nil {
        return fmt.Errorf("failed to update transaction status: %v", err)
    }
    
    rows, _ := result.RowsAffected()
    log.Printf("✅ Transaction status updated: id=%d, status=%s, rows=%d", transactionID, status, rows)
    
    return nil
}

// GetPendingTransactions возвращает pending транзакции для recovery
func (r *Repository) GetPendingTransactions(ctx context.Context) ([]Transaction, error) {
    rows, err := r.db.QueryContext(ctx,
        `SELECT id, event_nonce, user_address, amount, source_chain_id, target_chain_id, target_address, 
                status, signature, created_at, updated_at
         FROM transactions 
         WHERE status = $1`,
        StatePending)
    
    if err != nil {
        return nil, fmt.Errorf("failed to query pending transactions: %v", err)
    }
    defer rows.Close()
    
    var transactions []Transaction
    for rows.Next() {
        var t Transaction
        err := rows.Scan(&t.ID, &t.EventNonce, &t.UserAddress, &t.Amount, &t.SourceChainID, &t.TargetChainID,
            &t.TargetAddress, &t.Status, &t.Signature, &t.CreatedAt, &t.UpdatedAt)
        if err != nil {
            return nil, err
        }
        transactions = append(transactions, t)
    }
    
    return transactions, nil
}

// CheckNonce проверяет и обновляет nonce для защиты от replay атак
func (r *Repository) CheckNonce(ctx context.Context, userAddress string, chainID, nonce int64) (bool, error) {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return false, fmt.Errorf("failed to begin transaction: %v", err)
    }
    defer tx.Rollback()
    
    var lastNonce int64
    err = tx.QueryRowContext(ctx,
        `SELECT last_nonce FROM nonce_tracking WHERE user_address = $1 AND chain_id = $2`,
        userAddress, chainID).Scan(&lastNonce)
    
    if err == sql.ErrNoRows {
        // Первый nonce для этого пользователя и сети
        _, err = tx.ExecContext(ctx,
            `INSERT INTO nonce_tracking (user_address, chain_id, last_nonce, created_at, updated_at)
             VALUES ($1, $2, $3, NOW(), NOW())`,
            userAddress, chainID, nonce)
        if err != nil {
            return false, fmt.Errorf("failed to insert nonce tracking: %v", err)
        }
    } else if err != nil {
        return false, fmt.Errorf("failed to query nonce tracking: %v", err)
    } else {
        // Проверяем что nonce больше предыдущего
        if nonce <= lastNonce {
            log.Printf("⚠️ Replay attack detected: user=%s, chain=%d, nonce=%d (last=%d)", 
                userAddress, chainID, nonce, lastNonce)
            return false, nil
        }
        
        // Обновляем nonce
        _, err = tx.ExecContext(ctx,
            `UPDATE nonce_tracking SET last_nonce = $1, updated_at = NOW() 
             WHERE user_address = $2 AND chain_id = $3`,
            nonce, userAddress, chainID)
        if err != nil {
            return false, fmt.Errorf("failed to update nonce tracking: %v", err)
        }
    }
    
    if err := tx.Commit(); err != nil {
        return false, fmt.Errorf("failed to commit transaction: %v", err)
    }
    
    log.Printf("✅ Nonce checked: user=%s, chain=%d, nonce=%d", userAddress, chainID, nonce)
    return true, nil
}

// HealthCheck проверяет подключение к БД
func (r *Repository) HealthCheck(ctx context.Context) error {
    return r.db.PingContext(ctx)
}

// Close закрывает подключение к БД
func (r *Repository) Close() error {
    return r.db.Close()
}
