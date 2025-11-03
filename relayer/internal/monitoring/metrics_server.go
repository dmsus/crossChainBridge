package monitoring

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    log "github.com/sirupsen/logrus"
)

// StartMetricsServer запускает HTTP сервер для метрик Prometheus
func StartMetricsServer(port string) {
    http.Handle("/metrics", promhttp.Handler())
    go func() {
        server := &http.Server{
            Addr:    ":" + port,
            Handler: nil,
        }
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Errorf("Metrics server error: %v", err)
        }
    }()
    log.Infof("Metrics server started on :%s", port)
}
