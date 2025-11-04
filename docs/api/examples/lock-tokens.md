# 🔒 Lock Tokens - API Examples

## Basic Token Lock

### Request
```bash
curl -X POST http://localhost:8080/api/bridge/lock \
  -H "Content-Type: application/json" \
  -d '{
    "sourceNetwork": "ethereum",
    "targetNetwork": "polygon",
    "amount": "1000000000000000000",  # 1 ETH in wei
    "targetAddress": "0x29486B1429030420f00b9202dae6b76bd1b753c4"
  }'
```
## Response
```json
{
  "transactionHash": "0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
  "status": "pending",
  "estimatedCompletionTime": "2024-01-15T10:35:00Z"
}
```
## Lock with Custom Token
## Request
```bash
curl -X POST http://localhost:8080/api/bridge/lock \
  -H "Content-Type: application/json" \
  -d '{
    "sourceNetwork": "ethereum",
    "targetNetwork": "polygon",
    "amount": "50000000000000000000",  # 50 USDC (6 decimals)
    "targetAddress": "0x29486B1429030420f00b9202dae6b76bd1b753c4",
    "tokenAddress": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
  }'
```
## Polygon to Ethereum Transfer
## Request
```bash
curl -X POST http://localhost:8080/api/bridge/lock \
  -H "Content-Type: application/json" \
  -d '{
    "sourceNetwork": "polygon",
    "targetNetwork": "ethereum",
    "amount": "5000000000000000000",  # 5 MATIC in wei
    "targetAddress": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
  }'
```
## Error Examples
## Insufficient Balance
```json
{
  "error": "INSUFFICIENT_BALANCE",
  "message": "User has insufficient token balance",
  "code": 400,
  "details": {
    "required": "1000000000000000000",
    "available": "500000000000000000"
  }
}
```
## Invalid Address
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
## Network Unavailable
```json
{
  "error": "NETWORK_UNAVAILABLE",
  "message": "Polygon network is currently unavailable",
  "code": 503,
  "details": {
    "network": "polygon",
    "estimatedRecovery": "2024-01-15T11:00:00Z"
  }
}
```