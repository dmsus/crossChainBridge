package monitoring

import (
    "time"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    // Бизнес метрики
    TransactionsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "bridge_transactions_total",
        Help: "Total number of bridge transactions",
    }, []string{"status", "source_chain", "target_chain"})

    TransactionDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
        Name:    "bridge_transaction_duration_seconds",
        Help:    "Duration of bridge transaction processing",
        Buckets: []float64{0.1, 0.5, 1, 5, 10, 30},
    }, []string{"status"})

    GasUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
        Name: "bridge_gas_used",
        Help: "Gas used for bridge transactions",
    }, []string{"chain", "function"})

    // Системные метрики
    EventsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "bridge_events_processed_total",
        Help: "Total number of events processed",
    }, []string{"type", "status"})

    DatabaseOperations = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "bridge_database_operations_total", 
        Help: "Total database operations",
    }, []string{"operation", "table"})

    // Метрики ошибок
    ErrorsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
        Name: "bridge_errors_total",
        Help: "Total number of errors",
    }, []string{"component", "type"})

    // Метрики блокчейн соединений
    BlockchainConnections = promauto.NewGaugeVec(prometheus.GaugeOpts{
        Name: "bridge_blockchain_connections",
        Help: "Current blockchain connections status",
    }, []string{"chain", "endpoint"})
)

// Хелпер функции для записи метрик
func RecordTransaction(startTime time.Time, status, sourceChain, targetChain string) {
    duration := time.Since(startTime).Seconds()
    TransactionsTotal.WithLabelValues(status, sourceChain, targetChain).Inc()
    TransactionDuration.WithLabelValues(status).Observe(duration)
}

func RecordGasUsed(chain, function string, gas uint64) {
    GasUsed.WithLabelValues(chain, function).Set(float64(gas))
}

func RecordEvent(eventType, status string) {
    EventsProcessed.WithLabelValues(eventType, status).Inc()
}

func RecordError(component, errorType string) {
    ErrorsTotal.WithLabelValues(component, errorType).Inc()
}

func RecordDatabaseOperation(operation, table string) {
    DatabaseOperations.WithLabelValues(operation, table).Inc()
}

func UpdateBlockchainConnection(chain, endpoint string, connected bool) {
    value := 0.0
    if connected {
        value = 1.0
    }
    BlockchainConnections.WithLabelValues(chain, endpoint).Set(value)
}
