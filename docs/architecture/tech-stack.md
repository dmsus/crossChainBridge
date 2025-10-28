# Technology Stack

## Core Development

### Backend (Relay Service)
- **Go 1.21+** - High-performance, concurrent processing
- **Ethereum Go SDK** - Native blockchain interactions  
- **PostgreSQL 14+** - ACID compliance for financial transactions

### Smart Contracts
- **Solidity 0.8.x** - Security-focused with overflow protection
- **Hardhat** - Development environment with testing
- **OpenZeppelin** - Audited contract libraries

## Infrastructure

### Containerization
- **Docker** - Consistent development environments
- **Docker Compose** - Local development stack

### Monitoring & Observability
- **Prometheus** - Metrics collection
- **Grafana** - Visualization
- **Structured Logging** - JSON-formatted logs

### Optional/Advanced
- **Kubernetes (Minikube)** - Container orchestration
- **Make** - Build automation

## Justification

### Why Go?
- **Performance** - Native compilation and efficient concurrency
- **Blockchain Ecosystem** - First-class Ethereum support  
- **Safety** - Strong typing and memory safety
- **Operations** - Single binary deployment

### Why PostgreSQL?
- **ACID Compliance** - Critical for financial transactions
- **Idempotency** - Reliable duplicate detection
- **JSON Support** - Flexible event storage
- **Production Ready** - Enterprise-grade reliability

### Why Hardhat?
- **Developer Experience** - Excellent testing and debugging
- **TypeScript Support** - Type-safe contract development
- **Plugin Ecosystem** - Rich tooling integration
- **Community** - Widely adopted and supported