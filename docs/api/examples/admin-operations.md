# ⚙️ Admin Operations - API Examples

**Note:** All admin endpoints require `X-API-Key` header.

## List Transactions with Filters

### Request
```bash
curl -H "X-API-Key: admin-key-12345" \
  "http://localhost:8080/api/admin/transactions?status=failed&network=ethereum&page=1&limit=20"
```
## Response
```json
{
  "transactions": [
    {
      "transactionHash": "0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
      "status": "failed",
      "sourceNetwork": "ethereum",
      "targetNetwork": "polygon",
      "amount": "1000000000000000000",
      "timestamp": "2024-01-15T10:30:00Z",
      "confirmations": 12,
      "bridgeFee": "50000000000000000",
      "retryCount": 2,
      "lastError": "INSUFFICIENT_GAS: Transaction ran out of gas",
      "internalId": "tx_abc123"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 5,
    "pages": 1
  }
}
```
## Retry Failed Transaction
## Request
```bash
curl -X POST -H "X-API-Key: admin-key-12345" \
  http://localhost:8080/api/admin/transactions/tx_abc123/retry
```
## Response
```json
{
  "message": "Retry initiated successfully",
  "transactionId": "tx_abc123",
  "newStatus": "pending"
}
```
## Emergency Pause Bridge
## Request
```bash
curl -X POST -H "X-API-Key: admin-key-12345" \
  -H "Content-Type: application/json" \
  -d '{"reason": "Emergency maintenance - critical vulnerability patch"}' \
  http://localhost:8080/api/admin/pause
```
## Response
```json
{
  "message": "Bridge paused successfully",
  "pausedAt": "2024-01-15T11:00:00Z",
  "reason": "Emergency maintenance - critical vulnerability patch"
}
```
## Get System Metrics
## Request
```bash
curl -H "X-API-Key: admin-key-12345" \
  http://localhost:8080/api/admin/metrics
```
## Response
```json
{
  "system": {
    "uptime": "72h15m30s",
    "version": "1.0.0",
    "environment": "production"
  },
  "transactions": {
    "total": 1500,
    "completed": 1420,
    "failed": 65,
    "pending": 15,
    "successRate": "94.67%"
  },
  "performance": {
    "averageProcessingTime": "2m30s",
    "peakTps": 45,
    "currentTps": 12
  },
  "networks": {
    "ethereum": {
      "status": "healthy",
      "lastBlock": 18943221,
      "latency": "120ms"
    },
    "polygon": {
      "status": "healthy", 
      "lastBlock": 45321876,
      "latency": "85ms"
    }
  }
}
```
## Error Responses
## Missing API Key
```json
{
  "error": "UNAUTHORIZED",
  "message": "API key required for admin endpoints",
  "code": 401
}
```
## Invalid API Key
```json
{
  "error": "FORBIDDEN",
  "message": "Invalid API key",
  "code": 403
}
```
## Transaction Not Found (Retry)
```json
{
  "error": "TRANSACTION_NOT_FOUND",
  "message": "Transaction with id tx_abc123 not found",
  "code": 404
}
```
## Rate Limiting
Admin endpoints have higher rate limits:

- 1000 requests per minute for admin endpoints

- 100 requests per minute for public endpoints

## Rate Limit Exceeded
```json
{
  "error": "RATE_LIMIT_EXCEEDED",
  "message": "Rate limit exceeded for admin endpoints",
  "code": 429,
  "details": {
    "limit": 1000,
    "remaining": 0,
    "resetIn": 30
  }
}
```