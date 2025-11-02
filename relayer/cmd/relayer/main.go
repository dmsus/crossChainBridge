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
    "github.com/dmsus/crossChainBridge/relayer/internal/processor"
    "github.com/dmsus/crossChainBridge/relayer/internal/sender"
    "github.com/dmsus/crossChainBridge/relayer/pkg/database"
)

func main() {
    log.Println("🚀 Starting Cross-Chain Bridge Relayer with Idempotency")

    cfg, err := config.Load("staging")
    if err != nil {
        log.Fatalf("❌ Failed to load config: %v", err)
    }

    // Создаем подключение к БД
    dbRepo, err := database.SetupDatabase(database.Config{
        Host:     cfg.Database.Host,
        Port:     cfg.Database.Port,
        User:     cfg.Database.User,
        Password: cfg.Database.Password,
        DBName:   cfg.Database.Name,
        SSLMode:  "disable",
    })
    if err != nil {
        log.Fatalf("❌ Failed to setup database: %v", err)
    }
    defer dbRepo.Close()

    // Проверяем здоровье БД
    if err := dbRepo.HealthCheck(context.Background()); err != nil {
        log.Fatalf("❌ Database health check failed: %v", err)
    }
    log.Println("✅ Database health check passed")

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

    // Создаем процессор с идемпотентностью
    bridgeProcessor := processor.NewBridgeProcessor(processor.Config{
        PolygonSender: polygonSender,
        Repository:    dbRepo,
        MaxRetries:    3,
    })

    // Восстанавливаем pending транзакции при запуске
    if err := bridgeProcessor.RecoverPendingTransactions(context.Background()); err != nil {
        log.Printf("⚠️ Failed to recover pending transactions: %v", err)
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

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Запускаем обработку событий с идемпотентностью
    go processEventsWithIdempotency(ctx, ethListener, bridgeProcessor)

    // Запускаем listener
    if err := ethListener.Start(ctx); err != nil {
        log.Fatalf("❌ Failed to start Ethereum listener: %v", err)
    }

    log.Println("✅ Relayer with idempotency started successfully. Waiting for events...")

    // Ожидаем сигнал завершения
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    log.Println("🛑 Shutting down relayer...")
    cancel()
}

func processEventsWithIdempotency(ctx context.Context, listener *eventlistener.EthereumListener, processor *processor.BridgeProcessor) {
    for {
        select {
        case event := <-listener.Events():
            log.Printf("📦 Received event: user=%s, amount=%s, nonce=%s, targetChain=%s",
                event.User.Hex(), event.Amount.String(), event.Nonce.String(), event.TargetChainID.String())
            
            // Проверяем, что это перевод в Polygon (chainID 80002)
            if event.TargetChainID.Uint64() == 80002 {
                log.Println("🎯 This event is for Polygon network, processing with idempotency...")
                
                // Обрабатываем событие с гарантией идемпотентности - передаем указатель!
                if err := processor.ProcessEvent(ctx, &event); err != nil {
                    log.Printf("❌ Failed to process event: %v", err)
                } else {
                    log.Printf("✅ Event processed successfully with idempotency")
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
