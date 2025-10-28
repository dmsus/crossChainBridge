# Bridge API Documentation

## Overview
REST API for cross-chain token bridge operations, monitoring, and administration.

## Quick Start

### Base URLs
- **Development**: `http://localhost:8080`
- **Production**: `https://api.bridge.com`

### Authentication
Currently, the API uses API key authentication. Include your API key in the header:

X-API-Key: your-api-key-here


## Key Endpoints

### Bridge Operations
#### 1. Initiate Cross-Chain Transfer
POST /api/bridge/lock
Content-Type: application/json
```json
{
"sourceNetwork": "ethereum",
"targetNetwork": "polygon",
"amount": "1000000000000000000",
"targetAddress": "0x742E6F70B07533E0455c2e1A588aBc66a76b2f81"
}
```

#### 2. Check Transaction Status
GET /api/bridge/status/0x1234567890abcdef...

#### 3. Get Bridge Fees
GET /api/bridge/fees

#### 4. Get Transfer Limits
GET /api/bridge/limits

#### 5. Get Supported Tokens
GET /api/bridge/tokens

#### 6. Estimate Transfer
GET /api/bridge/estimate/1000000000000000000?sourceNetwork=ethereum&targetNetwork=polygon

### Monitoring
#### 7. Service Health
GET /health

#### 8. System Status
GET /api/system/status

#### 9. List Transactions
GET /api/transactions?page=1&limit=50

#### 10. Prometheus Metrics
GET /metrics

### Admin Operations
#### 11. List Transactions with Filters
GET /api/admin/transactions?status=failed&network=ethereum

#### 12. Retry Failed Transaction
POST /api/admin/transactions/{id}/retry

#### 13. Emergency Pause Bridge
POST /api/admin/pause
Content-Type: application/json
```json
{
"reason": "Emergency maintenance"
}
```

## OpenAPI Specification
The full API specification is available in [openapi.yaml](./openapi.yaml).

You can view interactive documentation using:
- [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [Redoc](https://redoc.ly/)

## Error Handling
All errors follow the standard format:
```json
{
  "error": "ValidationError",
  "message": "Invalid target address format",
  "code": 400,
  "details": {
    "field": "targetAddress"
  }
}
```
## Rate Limiting
**Bridge Operations**: 10 requests per minute per IP

**Monitoring**: 60 requests per minute per IP

**Admin**: 5 requests per minute per API key

**System**: 30 requests per minute per IP

## Response Examples
### Successful Lock
```json
{
  "transactionHash": "0x123...",
  "status": "pending",
  "estimatedCompletionTime": "2023-12-01T10:30:00Z"
}
```
### Transaction Status
```json
{
  "transactionHash": "0x123...",
  "status": "processing",
  "sourceNetwork": "ethereum",
  "targetNetwork": "polygon",
  "amount": "1000000000000000000",
  "timestamp": "2023-12-01T10:00:00Z",
  "confirmations": 12,
  "bridgeFee": "10000000000000000"
}
```
### System Status
```json
{
  "overall": "healthy",
  "components": {
    "database": {
      "status": "healthy",
      "latency": 45,
      "lastCheck": "2023-12-01T10:29:30Z"
    },
    "ethereumNode": {
      "status": "healthy", 
      "latency": 120,
      "lastCheck": "2023-12-01T10:29:25Z"
    }
  }
}
```
### Testing
Use the development server for testing:
```bash
curl -X GET http://localhost:8080/health
curl -X GET http://localhost:8080/api/bridge/fees
```
