# Cross-Chain Bridge Relayer

Go service for cross-chain bridge between Ethereum Sepolia and Polygon Amoy.

## Project Structure
relayer/
├── cmd/relayer/main.go # Application entry point
├── internal/ # Private application code
│ ├── config/ # Configuration management
│ ├── ethclient/ # Blockchain clients
│ ├── eventlistener/ # Event subscription
│ ├── signer/ # Cryptographic signing
│ ├── processor/ # Business logic
│ └── monitoring/ # Metrics and health checks
├── pkg/ # Public library code
│ ├── database/ # PostgreSQL layer
│ └── logger/ # Structured logging
├── api/ # API definitions
├── migrations/ # Database migrations
└── docs/ # Documentation


## Quick Start

```bash
# Local development with Anvil
go run cmd/relayer/main.go local

# Staging environment  
go run cmd/relayer/main.go staging
```
## Configuration
.env.local - Local development with Anvil

.env.staging - Testnet environments

## Dependencies
github.com/ethereum/go-ethereum - Blockchain interaction

github.com/lib/pq - PostgreSQL driver

github.com/spf13/viper - Configuration management

github.com/sirupsen/logrus - Structured logging

github.com/prometheus/client_golang - Metrics and monitoring
