# 📊 Check Transaction Status - API Examples

## Get Transaction Status

### Request
```bash
curl http://localhost:8080/api/bridge/status/0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8
```
## Status Response Examples
## Pending Status
```json
{
  "transactionHash": "0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
  "status": "pending",
  "sourceNetwork": "ethereum",
  "targetNetwork": "polygon", 
  "amount": "1000000000000000000",
  "timestamp": "2024-01-15T10:30:00Z",
  "confirmations": 2,
  "bridgeFee": "50000000000000000"
}
```
## Processing Status
```json
{
  "transactionHash": "0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
  "status": "processing",
  "sourceNetwork": "ethereum",
  "targetNetwork": "polygon",
  "amount": "1000000000000000000", 
  "timestamp": "2024-01-15T10:30:00Z",
  "confirmations": 12,
  "bridgeFee": "50000000000000000",
  "targetTransactionHash": "0x9ebb4f17f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9"
}
```
## Completed Status
```json
{
  "transactionHash": "0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
  "status": "completed",
  "sourceNetwork": "ethereum", 
  "targetNetwork": "polygon",
  "amount": "1000000000000000000",
  "timestamp": "2024-01-15T10:30:00Z",
  "confirmations": 64,
  "bridgeFee": "50000000000000000",
  "targetTransactionHash": "0x9ebb4f17f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9",
  "completedAt": "2024-01-15T10:33:00Z"
}
```
## Failed Status
```json
{
  "transactionHash": "0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
  "status": "failed",
  "sourceNetwork": "ethereum",
  "targetNetwork": "polygon",
  "amount": "1000000000000000000",
  "timestamp": "2024-01-15T10:30:00Z", 
  "confirmations": 12,
  "bridgeFee": "50000000000000000",
  "error": "INSUFFICIENT_GAS",
  "errorMessage": "Transaction ran out of gas",
  "failedAt": "2024-01-15T10:31:00Z"
}
```
## Error Responses
## Transaction Not Found
```json
{
  "error": "TRANSACTION_NOT_FOUND",
  "message": "Transaction with hash 0x123... not found",
  "code": 404
}
```
## Invalid Hash Format
```json
{
  "error": "VALIDATION_ERROR",
  "message": "Invalid transaction hash format",
  "code": 400,
  "details": {
    "field": "txHash", 
    "constraint": "must be valid transaction hash"
  }
}
```
## Monitoring Multiple Transactions
## Get Recent Transactions
```bash
curl "http://localhost:8080/api/transactions?page=1&limit=10"
```
## Response
```json
{
  "transactions": [
    {
      "transactionHash": "0x8daa4f16f8b8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
      "status": "completed",
      "sourceNetwork": "ethereum",
      "targetNetwork": "polygon",
      "amount": "1000000000000000000",
      "timestamp": "2024-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 150,
    "pages": 15
  }
}
```