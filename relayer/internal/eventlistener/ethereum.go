package eventlistener

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "math/big"
    "os"
    "time"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

type TokensLockedEvent struct {
    User          common.Address `json:"user"`
    Amount        *big.Int       `json:"amount"`
    Nonce         *big.Int       `json:"nonce"`
    TargetAddress common.Address `json:"targetAddress"`
    TargetChainID *big.Int       `json:"targetChainId"`
    TxHash        common.Hash    `json:"transactionHash"`
    BlockNumber   uint64         `json:"blockNumber"`
}

type Config struct {
    RPCEndpoint    string
    WSEndpoint     string
    ContractAddr   string
    ReconnectDelay time.Duration
    MaxRetries     int
}

type EthereumListener struct {
    client         *ethclient.Client
    wsClient       *ethclient.Client
    contract       common.Address
    contractABI    abi.ABI
    eventChan      chan TokensLockedEvent
    reconnectDelay time.Duration
    maxRetries     int
}

func NewEthereumListener(cfg Config) (*EthereumListener, error) {
    httpClient, err := ethclient.Dial(cfg.RPCEndpoint)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Ethereum RPC: %v", err)
    }

    wsClient, err := ethclient.Dial(cfg.WSEndpoint)
    if err != nil {
        httpClient.Close()
        return nil, fmt.Errorf("failed to connect to Ethereum WebSocket: %v", err)
    }

    contract := common.HexToAddress(cfg.ContractAddr)

    // ABI файл теперь в текущей папке
    abiPath := "contracts/out/BridgeEthereum.sol/BridgeEthereum.json"
    abiData, err := os.ReadFile(abiPath)
    if err != nil {
        httpClient.Close()
        wsClient.Close()
        return nil, fmt.Errorf("failed to read ABI file from %s: %v", abiPath, err)
    }

    log.Printf("📁 Loading ABI from: %s", abiPath)

    var abiStruct struct {
        ABI abi.ABI `json:"abi"`
    }
    if err := json.Unmarshal(abiData, &abiStruct); err != nil {
        httpClient.Close()
        wsClient.Close()
        return nil, fmt.Errorf("failed to parse ABI JSON: %v", err)
    }

    return &EthereumListener{
        client:         httpClient,
        wsClient:       wsClient,
        contract:       contract,
        contractABI:    abiStruct.ABI,
        eventChan:      make(chan TokensLockedEvent, 100),
        reconnectDelay: cfg.ReconnectDelay,
        maxRetries:     cfg.MaxRetries,
    }, nil
}

func (el *EthereumListener) Start(ctx context.Context) error {
    log.Println("🚀 Starting Ethereum event listener...")

    if err := el.checkConnection(ctx); err != nil {
        return fmt.Errorf("connection check failed: %v", err)
    }

    go el.subscribeToEvents(ctx)

    return nil
}

func (el *EthereumListener) checkConnection(ctx context.Context) error {
    block, err := el.client.BlockNumber(ctx)
    if err != nil {
        return err
    }
    log.Printf("✅ Connected to Ethereum. Current block: %d", block)

    code, err := el.client.CodeAt(ctx, el.contract, nil)
    if err != nil {
        return err
    }
    if len(code) == 0 {
        return fmt.Errorf("no contract code at address %s", el.contract.Hex())
    }
    log.Printf("✅ Contract verified at: %s", el.contract.Hex())

    return nil
}

func (el *EthereumListener) subscribeToEvents(ctx context.Context) {
    retryCount := 0

    for {
        select {
        case <-ctx.Done():
            log.Println("🛑 Event listener stopped by context")
            return
        default:
            err := el.startSubscription(ctx)
            if err != nil {
                retryCount++
                log.Printf("❌ Subscription error (attempt %d/%d): %v", retryCount, el.maxRetries, err)

                if retryCount >= el.maxRetries {
                    log.Fatalf("💥 Max retries exceeded. Stopping listener.")
                    return
                }

                delay := el.reconnectDelay * time.Duration(1<<uint(retryCount-1))
                log.Printf("⏳ Reconnecting in %v...", delay)
                time.Sleep(delay)
            } else {
                retryCount = 0
            }
        }
    }
}

func (el *EthereumListener) startSubscription(ctx context.Context) error {
    log.Println("📡 Subscribing to TokensLocked events...")

    query := ethereum.FilterQuery{
        Addresses: []common.Address{el.contract},
    }

    logs := make(chan types.Log)
    sub, err := el.wsClient.SubscribeFilterLogs(ctx, query, logs)
    if err != nil {
        return fmt.Errorf("failed to subscribe to logs: %v", err)
    }
    defer sub.Unsubscribe()

    log.Println("✅ Successfully subscribed to events")

    for {
        select {
        case err := <-sub.Err():
            return fmt.Errorf("subscription error: %v", err)

        case vLog := <-logs:
            event, err := el.parseTokensLockedEvent(vLog)
            if err != nil {
                continue
            }

            log.Printf("🎯 Received TokensLocked event: user=%s amount=%s nonce=%s targetChain=%s",
                event.User.Hex(), event.Amount.String(), event.Nonce.String(), event.TargetChainID.String())

            select {
            case el.eventChan <- *event:
                log.Println("✅ Event queued for processing")
            default:
                log.Println("⚠️ Event channel full, dropping event")
            }

        case <-ctx.Done():
            log.Println("🛑 Subscription stopped by context")
            return nil
        }
    }
}

func (el *EthereumListener) parseTokensLockedEvent(vLog types.Log) (*TokensLockedEvent, error) {
    event := struct {
        User          common.Address
        Amount        *big.Int
        Nonce         *big.Int
        TargetAddress common.Address
        TargetChainID *big.Int
    }{}

    err := el.contractABI.UnpackIntoInterface(&event, "TokensLocked", vLog.Data)
    if err != nil {
        return nil, fmt.Errorf("not a TokensLocked event: %v", err)
    }

    return &TokensLockedEvent{
        User:          event.User,
        Amount:        event.Amount,
        Nonce:         event.Nonce,
        TargetAddress: event.TargetAddress,
        TargetChainID: event.TargetChainID,
        TxHash:        vLog.TxHash,
        BlockNumber:   vLog.BlockNumber,
    }, nil
}

func (el *EthereumListener) Events() <-chan TokensLockedEvent {
    return el.eventChan
}

func (el *EthereumListener) Stop() {
    log.Println("🛑 Stopping Ethereum event listener...")

    if el.client != nil {
        el.client.Close()
    }
    if el.wsClient != nil {
        el.wsClient.Close()
    }

    close(el.eventChan)
    log.Println("✅ Ethereum event listener stopped")
}
