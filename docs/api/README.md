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

### 1. Initiate Cross-Chain Transfer
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
### 2. Check Transaction Status
GET /api/bridge/status/0x1234567890abcdef...

### 3. Service Health
GET /health

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
Rate Limiting
Bridge Operations: 10 requests per minute per IP

Monitoring: 60 requests per minute per IP

Admin: 5 requests per minute per API key