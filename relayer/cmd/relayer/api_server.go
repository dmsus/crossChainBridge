package main

import (
    "strconv"

    "github.com/dmsus/crossChainBridge/relayer/internal/api/routes"
    "github.com/dmsus/crossChainBridge/relayer/internal/config"
    "github.com/dmsus/crossChainBridge/relayer/pkg/database"
    "github.com/sirupsen/logrus"
)

// startAPIServerWithSecurity запускает API сервер с security features
func startAPIServerWithSecurity(cfg *config.Config, dbRepo *database.Repository, securityComponents *SecurityComponents, logger *logrus.Logger) {
    if !cfg.API.Enabled {
        logger.Info("🔕 API server is disabled in configuration")
        return
    }

    // API ключи для аутентификации (в продакшене из env variables)
    apiKeys := []string{
        "admin-key-123", // TODO: Заменить на реальные ключи из конфигурации
        "monitor-key-456",
    }

    // Создаем router
    router := routes.NewAPIRouter(
        dbRepo,
        securityComponents.RateLimiter,
        apiKeys,
    )

    port := strconv.Itoa(cfg.API.Port)
    
    logger.Infof("🚀 Starting API server on port %s", port)
    logger.Infof("📊 Available endpoints:")
    logger.Infof("   GET  /health                    - Health check")
    logger.Infof("   GET  /metrics                   - Prometheus metrics") 
    logger.Infof("   GET  /api/system/status         - System status")
    logger.Infof("   POST /api/bridge/lock           - Lock tokens")
    logger.Infof("   GET  /api/bridge/fees           - Bridge fees")
    logger.Infof("   GET  /api/bridge/limits         - Transfer limits")
    logger.Infof("   GET  /api/bridge/tokens         - Supported tokens")
    logger.Infof("   GET  /api/bridge/estimate/{amt} - Transfer estimate")
    logger.Infof("   GET  /api/bridge/status/{hash}  - Transaction status")
    logger.Infof("   GET  /api/transactions          - List transactions")
    logger.Infof("   GET  /api/admin/transactions    - Admin transactions (requires API key)")
    logger.Infof("   POST /api/admin/transactions/{id}/retry - Retry transaction")
    logger.Infof("   POST /api/admin/pause           - Pause bridge")
    logger.Infof("   GET  /api/admin/metrics         - System metrics")

    // Запускаем сервер
    if err := router.StartServer(port); err != nil {
        logger.Fatalf("❌ Failed to start API server: %v", err)
    }
}
