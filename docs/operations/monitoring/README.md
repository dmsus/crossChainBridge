# 📊 Monitoring & Observability

Complete monitoring guide for the cross-chain bridge system.

## 🎯 Key Metrics to Monitor

### Business Metrics
- **Transaction Success Rate**: Target > 99%
- **Average Processing Time**: Target < 3 minutes  
- **Bridge Volume**: Transactions per hour/day
- **Failed Transactions**: By error type

### Technical Metrics
- **Database Connections**: Active connections vs pool size
- **Blockchain Node Health**: Response time and error rate
- **Queue Depth**: Pending transactions in processing queue
- **Memory Usage**: Should stay below 80%

## 🔧 Monitoring Setup

### Prometheus Metrics

The bridge exposes metrics at `/metrics` endpoint:

```bash
# Check metrics endpoint
curl http://localhost:9090/metrics

# Example metrics:
bridge_transactions_total{status="completed"} 1420
bridge_transactions_total{status="failed"} 65
bridge_processing_duration_seconds{quantile="0.95"} 150
bridge_database_connections_active 12
bridge_blockchain_ethereum_last_block 18943221
```
## Health Checks
**Endpoint**: GET /health

```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z",
  "components": {
    "database": "healthy",
    "ethereumNode": "healthy",
    "polygonNode": "healthy", 
    "messageQueue": "healthy",
    "signingService": "healthy"
  }
}
```
## System Status
**Endpoint**: GET /api/system/status

```json
{
  "overall": "healthy",
  "components": {
    "database": {
      "status": "healthy",
      "latency": 15,
      "lastCheck": "2024-01-15T10:30:00Z"
    },
    "ethereumNode": {
      "status": "healthy", 
      "latency": 120,
      "lastCheck": "2024-01-15T10:30:00Z"
    },
    "polygonNode": {
      "status": "healthy",
      "latency": 85, 
      "lastCheck": "2024-01-15T10:30:00Z"
    }
  }
}
```
## 🚨 Alerting Rules
## Critical Alerts (Page Immediately)
```yaml
# Database down
- alert: DatabaseDown
  expr: bridge_database_connections_active == 0
  for: 1m
  labels:
    severity: critical
  annotations:
    summary: "Database is unreachable"

# Blockchain node down  
- alert: EthereumNodeDown
  expr: bridge_blockchain_ethereum_last_block - time() > 120
  for: 2m
  labels:
    severity: critical

# High failure rate
- alert: HighFailureRate
  expr: rate(bridge_transactions_total{status="failed"}[5m]) > 0.1
  for: 5m
  labels:
    severity: critical
```
## Warning Alerts
```yaml
# High processing time
- alert: HighProcessingTime
  expr: bridge_processing_duration_seconds{quantile="0.95"} > 300
  for: 10m
  labels:
    severity: warning

# Queue building up
- alert: QueueBacklog
  expr: bridge_queue_pending_transactions > 100
  for: 5m
  labels:
    severity: warning
```
## 📈 Dashboard Examples
## Grafana Dashboard Queries
## Transaction Success Rate:

```text
sum(rate(bridge_transactions_total{status="completed"}[5m])) / 
sum(rate(bridge_transactions_total[5m]))
```
## Average Processing Time:

```text
histogram_quantile(0.95, rate(bridge_processing_duration_seconds_bucket[5m]))
```
## Active Database Connections:

```text
bridge_database_connections_active
```
## Troubleshooting Common Issues
## High Processing Time
1. Check blockchain node latency

2. Verify database performance

3. Check for queue congestion

## Increasing Failure Rate
1. Review error logs for patterns

2. Check gas price fluctuations

3. Verify smart contract status

## Database Connection Issues
1. Check connection pool settings

2. Verify PostgreSQL is running

3. Check network connectivity
