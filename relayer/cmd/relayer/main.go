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
    "github.com/dmsus/crossChainBridge/relayer/internal/monitoring"
    "github.com/dmsus/crossChainBridge/relayer/internal/processor"
    "github.com/dmsus/crossChainBridge/relayer/internal/security"
    "github.com/dmsus/crossChainBridge/relayer/internal/sender"
    "github.com/dmsus/crossChainBridge/relayer/pkg/database"
    "github.com/sirupsen/logrus"
)

func main() {
    // Определяем окружение (по умолчанию staging)
    env := getEnvironment()
    
    log.Printf("🚀 Starting Cross-Chain Bridge Relayer with environment: %s", env)

    // Загружаем конфигурацию
    cfg, err := config.Load(env)
    if err != nil {
        log.Fatalf("❌ Failed to load config: %v", err)
    }

    // Настраиваем логирование
    logger := setupLogging(cfg)

    // Создаем контекст для graceful shutdown
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Запускаем health server
    go startHealthServer(ctx, "8080")

    // Запускаем сервер метрик Prometheus
    go monitoring.StartMetricsServer("9090")

    // Создаем подключение к БД
    dbRepo, err := database.SetupDatabase(database.Config{
        Host:     cfg.Database.Host,
        Port:     cfg.Database.Port,
        User:     cfg.Database.User,
        Password: cfg.Database.Password,
        DBName:   cfg.Database.Name,
        SSLMode:  cfg.Database.SSLMode,
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

    // ИНИЦИАЛИЗИРУЕМ SECURITY КОМПОНЕНТЫ
    securityComponents, err := setupSecurity(cfg, dbRepo, logger)
    if err != nil {
        log.Fatalf("❌ Failed to setup security components: %v", err)
    }
    log.Println("✅ Security components initialized")

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
        MaxRetries:    cfg.Processor.MaxRetries,
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
        ReconnectDelay: cfg.Processor.RetryDelay,
        MaxRetries:     cfg.Processor.MaxRetries,
    })
    if err != nil {
        log.Fatalf("❌ Failed to create Ethereum listener: %v", err)
    }
    defer ethListener.Stop()

    // Запускаем обработку событий с идемпотентностью И security
    go processEventsWithSecurity(ctx, ethListener, bridgeProcessor, securityComponents)

    // Запускаем API с security middleware (если API включено)
    if cfg.API.Enabled {
        go startAPIServerWithSecurity(cfg, dbRepo, securityComponents, logger)
    }

    // Запускаем listener
    if err := ethListener.Start(ctx); err != nil {
        log.Fatalf("❌ Failed to start Ethereum listener: %v", err)
    }

    log.Printf("✅ Relayer with security features started successfully in %s environment. Waiting for events...", env)

    // Ожидаем сигнал завершения
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    log.Println("🛑 Shutting down relayer...")
    cancel()
}

// setupSecurity инициализирует все security компоненты
func setupSecurity(cfg *config.Config, dbRepo *database.Repository, logger *logrus.Logger) (*security.SecurityComponents, error) {
    // Rate Limiter
    rateLimiter := security.NewRateLimiter(security.RateLimitConfig{
        DefaultRequestsPerMinute: cfg.Security.RequestsPerMinute,
        DefaultBurstSize:         cfg.Security.BurstSize,
        CleanupInterval:          1 * time.Minute,
    }, logger)

    // Replay Protector
    replayProtector := security.NewReplayProtector(dbRepo, logger)

    // Security Middleware
    securityConfig := security.SecurityConfig{
        EnableRateLimiting: cfg.Security.EnableRateLimiting,
        RequestsPerMinute:  cfg.Security.RequestsPerMinute,
        BurstSize:         cfg.Security.BurstSize,
        TimestampWindow:   cfg.Security.TimestampWindow,
        BlockedIPs:        cfg.Security.BlockedIPs,
    }
    middleware := security.NewSecurityMiddleware(rateLimiter, replayProtector, securityConfig, logger)

    return &security.SecurityComponents{
        RateLimiter:     rateLimiter,
        ReplayProtector: replayProtector,
        Middleware:      middleware,
    }, nil
}

func getEnvironment() string {
    if env := os.Getenv("APP_ENV"); env != "" {
        return env
    }
    return "staging" // default
}

func setupLogging(cfg *config.Config) *logrus.Logger {
    logger := logrus.New()
    
    // Устанавливаем уровень логирования
    level, err := logrus.ParseLevel(cfg.GetLogLevel())
    if err != nil {
        level = logrus.InfoLevel
    }
    logger.SetLevel(level)

    // Устанавливаем формат
    if cfg.Monitoring.LogFormat == "json" {
        logger.SetFormatter(&logrus.JSONFormatter{})
    } else {
        logger.SetFormatter(&logrus.TextFormatter{
            FullTimestamp: true,
        })
    }

    log.Printf("🔧 Log level: %s, format: %s", cfg.GetLogLevel(), cfg.Monitoring.LogFormat)
    return logger
}

func processEventsWithSecurity(ctx context.Context, listener *eventlistener.EthereumListener, processor *processor.BridgeProcessor, security *security.SecurityComponents) {
    for {
        select {
        case event := <-listener.Events():
            log.Printf("📦 Received event: user=%s, amount=%s, nonce=%s, targetChain=%s",
                event.User.Hex(), event.Amount.String(), event.Nonce.String(), event.TargetChainID.String())
            
            // ПРОВЕРЯЕМ REPLAY PROTECTION
            if security.ReplayProtector != nil {
                valid, err := security.ReplayProtector.ValidateNonceSequence(ctx, event.User.Hex(), event.TargetChainID.Int64(), event.Nonce.Int64())
                if err != nil {
                    log.Printf("❌ Replay protection check failed: %v", err)
                    continue
                }
                if !valid {
                    log.Printf("⚠️ Rejected duplicate or out-of-order nonce: %s", event.Nonce.String())
                    continue
                }
            }
            
            // Проверяем, что это перевод в Polygon (chainID 80002)
            if event.TargetChainID.Uint64() == 80002 {
                log.Println("🎯 This event is for Polygon network, processing with security...")
                
                // Обрабатываем событие с гарантией идемпотентности
                if err := processor.ProcessEvent(ctx, &event); err != nil {
                    log.Printf("❌ Failed to process event: %v", err)
                } else {
                    log.Printf("✅ Event processed successfully with security")
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
