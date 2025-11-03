# 🚀 CrossChainBridge: Ethereum Sepolia ↔ Polygon Amoy

**Production-ready cross-chain bridge implementation between Ethereum Sepolia and Polygon Amoy testnets.**

## 🎯 Project Status
**Phase 1 Completed** ✅ - Architecture and planning phase finished  
**Phase 2 Completed** ✅ - Smart contracts development finished  
**Phase 3 Completed** ✅ - Go relay service with idempotency finished  
**Phase 4 Starting** 🔜 - Infrastructure and testing


## 🛠 Technology Stack
- **Smart Contracts**: Solidity, Foundry, OpenZeppelin
- **Relay Service**: Go, PostgreSQL, Docker  
- **Blockchain**: Ethereum Sepolia (11155111), Polygon Amoy (80002)
- **Infrastructure**: Docker, Kubernetes, Prometheus
- **Security**: EIP-712 Signatures, Reentrancy Protection, Idempotency

## 📚 Documentation
- [Architecture](./docs/architecture/) - System design and C4 models
- [API Specification](./docs/api/) - REST API documentation  
- [Analysis](./docs/analysis/) - Requirements and user stories
- [Configuration](./relayer/docs/) - Environment-based config management
- [Smart Contracts](./contracts/) - Production-ready bridge contracts

## 🚀 Development Progress

### ✅ Phase 1: Analytics & Planning (COMPLETED)
- System architecture and C4 models
- Functional and non-functional requirements  
- User stories and use cases
- API specification (OpenAPI 3.0)
- Technology stack decisions

### ✅ Phase 2: Smart Contracts (COMPLETED)
- **BridgeEthereum.sol** - Token locking contract with event emission
- **BridgePolygon.sol** - Token release contract with EIP-712 signatures  
- **Token Contracts** - ERC20 tokens for both networks
- **Comprehensive Test Suite** - 13/13 tests passing (100% coverage)
- Security features: ReentrancyGuard, input validation, nonce tracking

### ✅ Phase 3: Go Relay Service (COMPLETED)
- **Ethereum Event Listener** - Real-time WebSocket subscription
- **Polygon Transaction Sender** - EIP-1559 transactions with gas optimization
- **EIP-712 Cryptographic Signatures** - Secure cross-chain verification
- **PostgreSQL Idempotency** - Guaranteed exactly-once processing
- **Configuration Management** - Environment-based config system

## 🏗️ System Architecture
Ethereum Sepolia → Event Listener → Idempotency Check → EIP-712 Signer → Polygon Amoy
✅ ✅ ✅ ✅ ✅


### 📊 Smart Contracts Overview

#### BridgeEthereum.sol
```solidity
function lockTokens(uint256 amount, address targetAddress, uint256 targetChainId)
```
- Locks tokens on Ethereum side

- Emits TokensLocked event for relayer

- Full input validation and security checks

#### BridgePolygon.sol
```solidity
function releaseTokens(address user, uint256 amount, uint256 nonce, bytes memory signature)
```
- Releases tokens on Polygon side

- EIP-712 signature verification

- Nonce-based idempotency protection

## Key Features
## Security
- EIP-712 Signatures - Cryptographic proof of legitimacy

- Reentrancy Protection - Guard against attack vectors

- Nonce Tracking - Replay attack prevention

- Input Validation - Comprehensive parameter checks

## Reliability
- Idempotent Processing - Exactly-once event handling

- Automatic Recovery - Pending transaction restoration

- Connection Resilience - Exponential backoff reconnection

- State Management - Full transaction lifecycle tracking

## Performance
- Event-Driven Architecture - Real-time processing

- Connection Pooling - Optimized database connections

- Gas Optimization - EIP-1559 transaction pricing

- Batch Processing - Efficient event handling

## Deployed Contracts
## Ethereum Sepolia
- **TokenEthereum**: 0x5CFdE9C777be47FC4a401c918181DD92BA4c81Cc

- **BridgeEthereum**: 0xc2766cBFc1Dc3E95058bf09c929E7D6226b10187

## Polygon Amoy
- **TokenPolygon**: 0x5CFdE9C777be47FC4a401c918181DD92BA4c81Cc

- **BridgePolygon**: 0xc2766cBFc1Dc3E95058bf09c929E7D6226b10187

## Testing
## Smart Contracts
```bash
cd contracts
forge test -vv
```
**Results**: 13/13 tests passing ✅

## Go Relay Service
```bash
cd relayer
go test ./...
```
## Quick Start
## Prerequisites
- Go 1.21+

- PostgreSQL 14+

- Foundry

- Docker

## Running the Relay Service
```bash
# Development environment
APP_ENV=dev go run cmd/relayer/main.go

# Staging environment (default)
go run cmd/relayer/main.go

# With environment variables
export ETH_PRIVATE_KEY="0x..."
export POLYGON_PRIVATE_KEY="0x..."
APP_ENV=staging go run cmd/relayer/main.go
```
## Environment Configuration
The system supports multiple environments:

- dev - Local development with Anvil

- staging - Test networks (Sepolia/Amoy)

- production - Mainnet deployment

## 🔜 Phase 4: Infrastructure & Testing (NEXT)
- Docker containerization

- Kubernetes deployment

- Integration testing

- E2E testing pipeline

- Monitoring and alerting

## 📈 Project Statistics
- 29 Tasks defined and tracked

- 19 Tasks completed (65%)

- 3 Phases fully implemented

- 4 Contracts deployed to testnets

- 100% Test Coverage for critical paths

## 🤝 Contributing
1. Fork the repository

2. Create a feature branch

3. Submit a pull request

## 📄 License
MIT License