# 🔌 Cross-Chain Bridge API Documentation

Complete REST API reference for the cross-chain bridge based on OpenAPI 3.0 specification.

## 📋 API Overview

| Category | Endpoints | Authentication |
|----------|-----------|----------------|
| **Bridge Operations** | Lock tokens, check status, get fees | Public |
| **Monitoring** | Health, metrics, transactions | Public |
| **Admin** | Transactions management, emergency controls | API Key Required |

## 🔐 Authentication

### Public Endpoints
No authentication required for bridge operations and monitoring.

### Admin Endpoints
Require API Key in header:
```bash
curl -H "X-API-Key: your-admin-key" http://localhost:8080/api/admin/transactions
```
## 🌐 Base URLs
- **Development**: http://localhost:8080

- **Production**: https://api.bridge.com

## 🚀 Quick Start Examples
## 1. Check System Health
```bash
curl http://localhost:8080/health
```
## Response:

```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z",
  "components": {
    "database": "healthy",
    "ethereumNode": "healthy", 
    "polygonNode": "healthy",
    "messageQueue": "healthy"
  }
}
```
## 2. Lock Tokens (Initiate Transfer)
```bash
curl -X POST http://localhost:8080/api/bridge/lock \
  -H "Content-Type: application/json" \
  -d '{
    "sourceNetwork": "ethereum",
    "targetNetwork": "polygon", 
    "amount": "1000000000000000000",
    "targetAddress": "0x29486B1429030420f00b9202dae6b76bd1b753c4"
  }'
```
## Response:

```json
{
  "transactionHash": "0x123...abc",
  "status": "pending",
  "estimatedCompletionTime": "2024-01-15T10:35:00Z"
}
```
## 3. Check Transaction Status
```bash
curl http://localhost:8080/api/bridge/status/0x123...abc
```
## Response:

```json
{
  "transactionHash": "0x123...abc",
  "status": "processing",
  "sourceNetwork": "ethereum",
  "targetNetwork": "polygon",
  "amount": "1000000000000000000",
  "timestamp": "2024-01-15T10:30:00Z",
  "confirmations": 12,
  "bridgeFee": "50000000000000000"
}
```
## 📊 Bridge Operations
## POST /api/bridge/lock
Initiate token lock for cross-chain transfer.

## Request Body:

```json
{
  "sourceNetwork": "ethereum",
  "targetNetwork": "polygon", 
  "amount": "1000000000000000000",
  "targetAddress": "0x29486B1429030420f00b9202dae6b76bd1b753c4",
  "tokenAddress": "0x5CFdE9C777be47FC4a401c918181DD92BA4c81Cc"
}
```
## GET /api/bridge/status/{txHash}
Get transaction status by hash.

## GET /api/bridge/fees
Get current bridge fee structure.

```bash
curl http://localhost:8080/api/bridge/fees
```
## GET /api/bridge/limits
Get transfer limits.

```bash
curl http://localhost:8080/api/bridge/limits
```
## GET /api/bridge/tokens
Get supported tokens.

```bash
curl http://localhost:8080/api/bridge/tokens
```
## GET /api/bridge/estimate/{amount}
Estimate transfer time and fees.

```bash
curl "http://localhost:8080/api/bridge/estimate/1000000000000000000?sourceNetwork=ethereum&targetNetwork=polygon"
```
## 📈 Monitoring Endpoints
## GET /health
System health check.

## GET /metrics
Prometheus metrics (text format).

## GET /api/system/status
Detailed system component status.

## GET /api/transactions
List recent transactions with pagination.

```bash
curl "http://localhost:8080/api/transactions?page=1&limit=50"
```
## ⚙️ Admin Endpoints
## GET /api/admin/transactions
List transactions with filters (requires API key).

```bash
curl -H "X-API-Key: admin-key" \
  "http://localhost:8080/api/admin/transactions?status=failed&network=ethereum&page=1&limit=50"
```
## POST /api/admin/transactions/{id}/retry
Retry failed transaction.

```bash
curl -X POST -H "X-API-Key: admin-key" \
  http://localhost:8080/api/admin/transactions/123/retry
```
## POST /api/admin/pause
Emergency pause bridge.

```bash
curl -X POST -H "X-API-Key: admin-key" \
  -H "Content-Type: application/json" \
  -d '{"reason": "emergency maintenance"}' \
  http://localhost:8080/api/admin/pause
```
## 🚨 Error Responses
All errors follow this format:

```json
{
  "error": "VALIDATION_ERROR",
  "message": "Invalid target address format",
  "code": 400,
  "details": {
    "field": "targetAddress",
    "constraint": "must be valid Ethereum address"
  }
}
```
## Common HTTP Status Codes:

- 200 - Success

- 400 - Bad Request (invalid parameters)

- 401 - Unauthorized (missing/invalid API key)

- 404 - Not Found

- 429 - Rate Limit Exceeded

- 500 - Internal Server Error

- 503 - Service Unavailable

## 📝 Notes
- **Amount Format**: Always use smallest unit (wei for ETH)

- **Address Format**: Must be valid Ethereum addresses (0x...)

- **Rate Limits**: 100 requests/minute for public endpoints

- **Idempotency**: Lock operations are idempotent

## 🔗 OpenAPI Specification
Full API specification available at: /docs/api/openapi.yaml
