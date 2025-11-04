# 🔧 Troubleshooting Guide

Common issues and solutions for the cross-chain bridge system.

## 🚨 Critical Issues

### Bridge Not Processing Transactions

**Symptoms**:
- Transactions stuck in "pending" status
- No new transactions being processed
- Error logs showing processing failures

**Diagnosis**:
```bash
# Check relayer status
curl http://localhost:8081/health

# Check database connections
docker-compose logs postgres | tail -20

# Check blockchain connections  
curl -X POST $ETH_RPC_URL -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```
## Solutions:

## 1. Restart relayer service:

```bash
docker-compose restart relayer
```
## 2. Check private key balance:

```bash
cast balance $RELAYER_ADDRESS --rpc-url $ETH_RPC_URL
```
## 3. Reset stuck transactions:

```sql
UPDATE transactions SET status = 'pending' 
WHERE status = 'processing' AND updated_at < NOW() - INTERVAL '10 minutes';
```
## Database Connection Issues
**Symptoms**:

- Health check shows database unhealthy

- "connection refused" errors in logs

- API endpoints returning 503 errors

## Diagnosis:

```bash
# Check PostgreSQL status
docker-compose ps postgres

# Check connection from relayer
docker-compose exec relayer nc -z postgres 5432

# Check database logs
docker-compose logs postgres
```
## Solutions:

## 1. Restart PostgreSQL:

```bash
docker-compose restart postgres
```
## 2. Check disk space:

```bash
docker-compose exec postgres df -h
```
## 3. Check connection limits:

```bash
docker-compose exec postgres psql -U bridge_user -d bridge_local -c "SELECT count(*) FROM pg_stat_activity;"
```
## Blockchain Node Issues
**Symptoms**:

- "connection timeout" errors

- Transactions failing with network errors

- Health check shows node unhealthy

**Diagnosis**:

```bash
# Test Ethereum connection
curl -X POST $ETH_RPC_URL -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}'

# Test Polygon connection  
curl -X POST $POLYGON_RPC_URL -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}'

# Check latest block
cast block latest --rpc-url $ETH_RPC_URL
```
## Solutions:

1. **Switch to backup RPC endpoint**:
Update environment variables to use backup URLs

2. **Check gas prices**:

```bash
cast gas-price --rpc-url $ETH_RPC_URL
```
3. **Verify chain ID**:

```bash
cast chain-id --rpc-url $ETH_RPC_URL
```
## ⚠️ Common Issues
## High Gas Prices
**Symptoms**:

- Transactions failing with "out of gas"

- High gas fees in transaction estimates

- Slow transaction confirmation

**Solutions**:

- Increase gas limit in configuration

- Wait for network congestion to decrease

- Use gas price estimation from multiple sources

## Nonce Conflicts
**Symptoms**:

- Transactions failing with "nonce too low"

- Multiple transactions with same nonce

- Transactions stuck in mempool

**Solutions**:

```bash
# Check current nonce
cast nonce $RELAYER_ADDRESS --rpc-url $ETH_RPC_URL

# Reset transaction pool (use with caution)
cast rpc --rpc-url $ETH_RPC_URL txpool_content
```
## Insufficient Balance
**Symptoms**:

- Transactions failing with "insufficient funds"

- Relayer address has low balance

- Bridge operations halted

**Solutions**:

1. **Fund relayer address**:

```bash
# Check balance
cast balance $RELAYER_ADDRESS --rpc-url $ETH_RPC_URL
```
2. **Emergency funding procedure**:

- Transfer funds to relayer address

- Verify balance updated

- Retry failed transactions

## 📊 Performance Issues
## Slow Transaction Processing
**Diagnosis**:

```bash
# Check processing queue
curl http://localhost:8081/api/system/status | jq '.components.messageQueue'

# Check database performance
docker-compose exec postgres psql -U bridge_user -d bridge_local -c "SELECT schemaname, tablename, seq_scan, seq_tup_read, idx_scan, idx_tup_fetch FROM pg_stat_user_tables;"
```
**Solutions**:

- Add database indexes (see performance optimization guide)

- Increase relayer resources

- Optimize blockchain RPC endpoints

## Memory Leaks
**Symptoms**:

- Relayer container restarting frequently

- Increasing memory usage over time

- "Out of memory" errors in logs

**Diagnosis**:

```bash
# Check memory usage
docker stats $(docker ps -q)

# Check Go garbage collection
curl http://localhost:9090/metrics | grep go_memstats
```
**Solutions**:

- Increase container memory limits

- Analyze memory profiles

- Update to latest Go version

## 🔍 Debugging Tools
## Log Analysis
**Key log patterns to monitor**:

```bash
# Transaction processing errors
docker-compose logs relayer | grep -i "error\|failed"

# Database connection issues
docker-compose logs relayer | grep -i "database\|postgres"

# Blockchain interactions
docker-compose logs relayer | grep -i "transaction\|gas\|nonce"
```
## Metrics Analysis
**Key metrics to check**:

```bash
# Success rate
curl http://localhost:9090/metrics | grep bridge_transactions_total

# Processing time
curl http://localhost:9090/metrics | grep bridge_processing_duration

# Resource usage
curl http://localhost:9090/metrics | grep -E "go_goroutines|go_memstats"
```
**Database Queries for Debugging**
```sql
-- Check transaction status distribution
SELECT status, COUNT(*) FROM transactions GROUP BY status;

-- Find stuck transactions
SELECT * FROM transactions 
WHERE status = 'processing' AND updated_at < NOW() - INTERVAL '5 minutes';

-- Check recent errors
SELECT * FROM transactions 
WHERE status = 'failed' AND created_at > NOW() - INTERVAL '1 hour'
ORDER BY created_at DESC;
```
## 🆘 Emergency Procedures
**Immediate Bridge Pause**
```bash
# Using admin API
curl -X POST -H "X-API-Key: $ADMIN_API_KEY" \
  http://localhost:8081/api/admin/pause \
  -d '{"reason": "Emergency maintenance"}'
```
**Service Rollback**
```bash
# Rollback to previous Docker image
docker-compose pull relayer:previous-version
docker-compose up -d
```
## Contact Support
**When to escalate**:

- Data corruption detected

- Security breach suspected

- Extended outage (> 1 hour)

- Financial impact to users

## Information to provide:

- Error logs and metrics

- Steps already attempted

- Impact assessment

- Timeline of events
