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
    "github.com/dmsus/crossChainBridge/relayer/internal/sender"
)

func main() {
    log.Println("🚀 Starting Cross-Chain Bridge Relayer")

    cfg, err := config.Load("staging")
    if err != nil {
        log.Fatalf("❌ Failed to load config: %v", err)
    }

    // Создаем Ethereum listener
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

    // Создаем Polygon sender
    polygonSender, err := sender.NewPolygonSender(sender.Config{
        RPCEndpoint:  cfg.Polygon.RPCURL,
        PrivateKey:   cfg.Polygon.PrivateKey,
        ContractAddr: cfg.Polygon.BridgeAddr,
    })
    if err != nil {
        log.Fatalf("❌ Failed to create Polygon sender: %v", err)
    }
    defer polygonSender.Close()

    // Проверяем здоровье Polygon подключения
    if err := polygonSender.HealthCheck(context.Background()); err != nil {
        log.Fatalf("❌ Polygon health check failed: %v", err)
    }
    log.Println("✅ Polygon sender health check passed")

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Запускаем обработку событий с интеграцией Polygon sender
    go processEvents(ctx, ethListener, polygonSender)

    // Запускаем listener
    if err := ethListener.Start(ctx); err != nil {
        log.Fatalf("❌ Failed to start Ethereum listener: %v", err)
    }

    log.Println("✅ Relayer started successfully. Waiting for events...")

    // Ожидаем сигнал завершения
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    log.Println("🛑 Shutting down relayer...")
    cancel()
}

func processEvents(ctx context.Context, listener *eventlistener.EthereumListener, polygonSender *sender.PolygonSender) {
    for {
        select {
        case event := <-listener.Events():
            log.Printf("📦 Processing event: user=%s, amount=%s, nonce=%s, targetChain=%s",
                event.User.Hex(), event.Amount.String(), event.Nonce.String(), event.TargetChainID.String())
            
            // Проверяем, что это перевод в Polygon (chainID 80002)
            if event.TargetChainID.Uint64() == 80002 {
                log.Println("🎯 This event is for Polygon network, processing...")
                
                // ВРЕМЕННО: используем заглушку для подписи
                // В #17 задаче реализуем настоящие EIP-712 подписи
                fakeSignature := []byte("fake_signature_for_testing")
                
                // Отправляем транзакцию в Polygon
                tx, err := polygonSender.SendReleaseTokens(ctx, event.User, event.Amount, event.Nonce, fakeSignature)
                if err != nil {
                    log.Printf("❌ Failed to send transaction to Polygon: %v", err)
                    // TODO: Добавить логику повторных попыток
                } else {
                    log.Printf("✅ Transaction sent to Polygon: %s", tx.Hash().Hex())
                }
            } else {
                log.Printf("⚠️ Skipping event for unknown chain: %s", event.TargetChainID.String())
            }
            
        case <-ctx.Done():
            log.Println("🛑 Event processor stopped")
            return
        }
    }
}
