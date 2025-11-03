package main

import (
    "context"
    "log"
    "net/http"
    "time"
)

func startHealthServer(ctx context.Context, port string) {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status": "healthy", "service": "cross-chain-bridge"}`))
    })
    
    mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status": "ready", "service": "cross-chain-bridge"}`))
    })
    
    server := &http.Server{
        Addr:    ":" + port,
        Handler: mux,
    }
    
    go func() {
        log.Printf("🏥 Health server starting on port %s", port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Printf("⚠️ Health server failed: %v", err)
        }
    }()
    
    // Graceful shutdown
    go func() {
        <-ctx.Done()
        log.Println("🛑 Shutting down health server...")
        
        shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        
        if err := server.Shutdown(shutdownCtx); err != nil {
            log.Printf("⚠️ Health server shutdown error: %v", err)
        } else {
            log.Println("✅ Health server stopped")
        }
    }()
}
