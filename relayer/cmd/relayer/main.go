package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/dmsus/crossChainBridge/relayer/internal/config"
    "github.com/dmsus/crossChainBridge/relayer/internal/eventlistener"
)

func main() {
    log.Println("🚀 Starting Cross-Chain Bridge Relayer")

    cfg, err := config.Load("staging")
    if err != nil {
        log.Fatalf("❌ Failed to load config: %v", err)
    }

    ethListener, err := eventlistener.NewEthereumListener(eventlistener.Config{
        RPCEndpoint:    cfg.Ethereum.RPCURL,
        WSEndpoint:     cfg.Ethereum.WsURL,
        ContractAddr:   cfg.Ethereum.BridgeAddr,
        ReconnectDelay: 5 * time.Second,
        MaxRetries:     10,
    })
    if err != nil {
        log.Fatalf("❌ Failed to create Ethereum listener: %v", err)
    }
    defer ethListener.Stop()

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go processEvents(ctx, ethListener)

    if err := ethListener.Start(ctx); err != nil {
        log.Fatalf("❌ Failed to start Ethereum listener: %v", err)
    }

    log.Println("✅ Relayer started successfully. Waiting for events...")

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    log.Println("🛑 Shutting down relayer...")
    cancel()
}

func processEvents(ctx context.Context, listener *eventlistener.EthereumListener) {
    for {
        select {
        case event := <-listener.Events():
            log.Printf("📦 Processing event: user=%s, amount=%s, nonce=%s, targetChain=%s",
                event.User.Hex(), event.Amount.String(), event.Nonce.String(), event.TargetChainID.String())
            
        case <-ctx.Done():
            log.Println("🛑 Event processor stopped")
            return
        }
    }
}
