package routes

import (
    "net/http"

    "github.com/dmsus/crossChainBridge/relayer/internal/api/handlers"
    "github.com/dmsus/crossChainBridge/relayer/internal/api/middleware"
    "github.com/dmsus/crossChainBridge/relayer/internal/security"
    "github.com/dmsus/crossChainBridge/relayer/pkg/database"
    "github.com/gorilla/mux"
)

// APIRouter настраивает все API routes
type APIRouter struct {
    router *mux.Router
}

// NewAPIRouter создает новый router с всеми endpoints
func NewAPIRouter(
    dbRepo *database.Repository,
    rateLimiter *security.RateLimiter,
    apiKeys []string,
) *APIRouter {
    r := mux.NewRouter()

    // Инициализируем handlers
    monitoringHandler := handlers.NewMonitoringHandler(dbRepo)
    adminHandler := handlers.NewAdminHandler(dbRepo)
    bridgeHandler := handlers.NewBridgeHandler()

    // Инициализируем middleware
    authMiddleware := middleware.NewAPIKeyAuth(apiKeys)
    securityMiddleware := middleware.NewSecurityMiddleware(rateLimiter)

    // Apply global middleware
    r.Use(securityMiddleware.RateLimitMiddleware)
    r.Use(middleware.SecurityHeadersMiddleware)
    r.Use(authMiddleware.Middleware)
    r.Use(middleware.AdminOnlyMiddleware)

    // Health and monitoring routes (public)
    r.HandleFunc("/health", monitoringHandler.HealthCheck).Methods("GET")
    r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        // TODO: Integrate with Prometheus metrics
        w.Write([]byte("# metrics endpoint - integrate with Prometheus\n"))
    }).Methods("GET")

    // System status routes (public)
    r.HandleFunc("/api/system/status", monitoringHandler.SystemStatus).Methods("GET")

    // Bridge operation routes (public)
    bridgeRouter := r.PathPrefix("/api/bridge").Subrouter()
    bridgeRouter.HandleFunc("/lock", bridgeHandler.LockTokens).Methods("POST")
    bridgeRouter.HandleFunc("/fees", bridgeHandler.GetFees).Methods("GET")
    bridgeRouter.HandleFunc("/limits", bridgeHandler.GetLimits).Methods("GET")
    bridgeRouter.HandleFunc("/tokens", bridgeHandler.GetSupportedTokens).Methods("GET")
    bridgeRouter.HandleFunc("/estimate/{amount}", bridgeHandler.EstimateTransfer).Methods("GET")
    bridgeRouter.HandleFunc("/status/{txHash}", monitoringHandler.GetTransactionStatus).Methods("GET")

    // Transaction routes (public)
    r.HandleFunc("/api/transactions", monitoringHandler.GetTransactions).Methods("GET")

    // Admin routes (protected)
    adminRouter := r.PathPrefix("/api/admin").Subrouter()
    adminRouter.HandleFunc("/transactions", adminHandler.GetAdminTransactions).Methods("GET")
    adminRouter.HandleFunc("/transactions/{id}/retry", adminHandler.RetryTransaction).Methods("POST")
    adminRouter.HandleFunc("/pause", adminHandler.PauseBridge).Methods("POST")
    adminRouter.HandleFunc("/metrics", adminHandler.GetSystemMetrics).Methods("GET")

    return &APIRouter{router: r}
}

// Handler возвращает HTTP handler
func (ar *APIRouter) Handler() http.Handler {
    return ar.router
}

// StartServer запускает HTTP сервер
func (ar *APIRouter) StartServer(port string) error {
    return http.ListenAndServe(":"+port, ar.router)
}
