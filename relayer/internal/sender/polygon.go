package sender

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "time"

    "github.com/dmsus/crossChainBridge/relayer/internal/signer"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

// PolygonSender отправляет транзакции в сеть Polygon Amoy
type PolygonSender struct {
    client     *ethclient.Client
    privateKey string
    contract   common.Address
    chainID    *big.Int
    signer     *signer.EIP712Signer
    abiEncoder *signer.ABIEncoder
}

// Config конфигурация для Polygon sender
type Config struct {
    RPCEndpoint  string
    PrivateKey   string
    ContractAddr string
}

// NewPolygonSender создает новый экземпляр отправителя
func NewPolygonSender(cfg Config) (*PolygonSender, error) {
    client, err := ethclient.Dial(cfg.RPCEndpoint)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Polygon RPC: %v", err)
    }

    // Получаем chainID сети
    chainID, err := client.ChainID(context.Background())
    if err != nil {
        client.Close()
        return nil, fmt.Errorf("failed to get chain ID: %v", err)
    }

    contract := common.HexToAddress(cfg.ContractAddr)

    // Создаем EIP712 signer
    eip712Signer, err := signer.NewEIP712Signer(signer.Config{
        PrivateKey:    cfg.PrivateKey,
        BridgeAddress: contract,
        ChainID:       chainID,
    })
    if err != nil {
        client.Close()
        return nil, fmt.Errorf("failed to create EIP712 signer: %v", err)
    }

    // Создаем ABI encoder
    abiEncoder := signer.NewABIEncoder()

    log.Printf("✅ Polygon sender initialized: chainID=%d, contract=%s", chainID, contract.Hex())

    return &PolygonSender{
        client:     client,
        privateKey: cfg.PrivateKey,
        contract:   contract,
        chainID:    chainID,
        signer:     eip712Signer,
        abiEncoder: abiEncoder,
    }, nil
}

// SendReleaseTokens отправляет транзакцию releaseTokens в Polygon
func (ps *PolygonSender) SendReleaseTokens(ctx context.Context, user common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
    log.Printf("🚀 Preparing releaseTokens transaction: user=%s, amount=%s, nonce=%s", 
        user.Hex(), amount.String(), nonce.String())

    // 1. Создаем EIP-712 подпись
    bridgeMessage := &signer.BridgeMessage{
        User:          user,
        Amount:        amount,
        Nonce:         nonce,
        TargetChainID: ps.chainID,
    }

    signature, err := ps.signer.SignBridgeMessage(bridgeMessage)
    if err != nil {
        return nil, fmt.Errorf("failed to sign bridge message: %v", err)
    }

    // 2. Создаем транзакцию с настоящей подписью
    tx, err := ps.createReleaseTokensTx(ctx, user, amount, nonce, signature)
    if err != nil {
        return nil, fmt.Errorf("failed to create transaction: %v", err)
    }

    // 3. Отправляем транзакцию
    err = ps.client.SendTransaction(ctx, tx)
    if err != nil {
        return nil, fmt.Errorf("failed to send transaction: %v", err)
    }

    log.Printf("✅ Transaction sent: hash=%s", tx.Hash().Hex())

    // 4. Мониторим подтверждение
    go ps.monitorTransaction(ctx, tx.Hash())

    return tx, nil
}

// createReleaseTokensTx создает подписанную транзакцию с ABI кодированием
func (ps *PolygonSender) createReleaseTokensTx(ctx context.Context, user common.Address, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
    // Парсим приватный ключ
    privateKey, err := crypto.HexToECDSA(ps.privateKey[2:])
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %v", err)
    }

    // Создаем transactor
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, ps.chainID)
    if err != nil {
        return nil, fmt.Errorf("failed to create transactor: %v", err)
    }

    // Устанавливаем параметры газа
    auth.GasLimit = uint64(300000)
    auth.Context = ctx

    // Получаем nonce для аккаунта
    accountNonce, err := ps.client.PendingNonceAt(ctx, auth.From)
    if err != nil {
        return nil, fmt.Errorf("failed to get nonce: %v", err)
    }

    // Получаем актуальные цены на газ (EIP-1559)
    gasTipCap, gasFeeCap, err := ps.getGasPrices(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to get gas prices: %v", err)
    }

    // Кодируем данные вызова через ABI
    data, err := ps.abiEncoder.EncodeReleaseTokensCall(user, amount, nonce, signature)
    if err != nil {
        return nil, fmt.Errorf("failed to encode ABI call: %v", err)
    }

    // Создаем динамическую транзакцию (EIP-1559)
    tx := types.NewTx(&types.DynamicFeeTx{
        ChainID:   ps.chainID,
        Nonce:     accountNonce,
        GasTipCap: gasTipCap,
        GasFeeCap: gasFeeCap,
        Gas:       auth.GasLimit,
        To:        &ps.contract,
        Value:     big.NewInt(0),
        Data:      data,
    })

    // Подписываем транзакцию
    signedTx, err := auth.Signer(auth.From, tx)
    if err != nil {
        return nil, fmt.Errorf("failed to sign transaction: %v", err)
    }

    return signedTx, nil
}

// getGasPrices получает актуальные цены на газ для EIP-1559
func (ps *PolygonSender) getGasPrices(ctx context.Context) (*big.Int, *big.Int, error) {
    header, err := ps.client.HeaderByNumber(ctx, nil)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to get header: %v", err)
    }

    gasTipCap := big.NewInt(1500000000) // 1.5 Gwei

    baseFee := header.BaseFee
    if baseFee == nil {
        baseFee = big.NewInt(10000000000) // fallback: 10 Gwei
    }

    gasFeeCap := new(big.Int).Add(
        new(big.Int).Mul(baseFee, big.NewInt(2)),
        gasTipCap,
    )

    log.Printf("⛽ Gas prices: tip=%s wei, feeCap=%s wei", 
        gasTipCap.String(), gasFeeCap.String())

    return gasTipCap, gasFeeCap, nil
}

// monitorTransaction отслеживает статус транзакции
func (ps *PolygonSender) monitorTransaction(ctx context.Context, txHash common.Hash) {
    log.Printf("👀 Monitoring transaction: %s", txHash.Hex())

    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            receipt, err := ps.client.TransactionReceipt(ctx, txHash)
            if err != nil {
                continue
            }

            if receipt != nil {
                if receipt.Status == 1 {
                    log.Printf("🎉 Transaction confirmed: %s (block %d)", 
                        txHash.Hex(), receipt.BlockNumber.Uint64())
                } else {
                    log.Printf("❌ Transaction failed: %s", txHash.Hex())
                }
                return
            }

        case <-ctx.Done():
            log.Printf("🛑 Transaction monitoring stopped: %s", txHash.Hex())
            return
        }
    }
}

// Close закрывает подключение
func (ps *PolygonSender) Close() {
    if ps.client != nil {
        ps.client.Close()
        log.Println("✅ Polygon sender closed")
    }
}

// HealthCheck проверяет подключение к Polygon
func (ps *PolygonSender) HealthCheck(ctx context.Context) error {
    _, err := ps.client.BlockNumber(ctx)
    if err != nil {
        return fmt.Errorf("Polygon connection failed: %v", err)
    }
    return nil
}

// VerifySignature проверяет EIP-712 подпись
func (ps *PolygonSender) VerifySignature(user common.Address, amount *big.Int, nonce *big.Int, signature []byte) (bool, common.Address, error) {
    bridgeMessage := &signer.BridgeMessage{
        User:          user,
        Amount:        amount,
        Nonce:         nonce,
        TargetChainID: ps.chainID,
    }
    return ps.signer.VerifySignature(bridgeMessage, signature)
}
