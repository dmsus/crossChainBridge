# CrossChainBridge: Ethereum Sepolia ↔ Polygon Amoy

Enterprise-grade cross-chain bridge implementation between Ethereum Sepolia and Polygon Amoy testnets.

## 🎯 Project Status
**Phase 1 Completed** ✅ - Architecture and planning phase finished  
**Phase 2 Completed** ✅ - Smart contracts development finished  
**Phase 3 Starting** 🚧 - Go relay service development

## 🛠 Technology Stack
- **Smart Contracts**: Solidity, Foundry, OpenZeppelin
- **Relay Service**: Go, PostgreSQL  
- **Infrastructure**: Docker, Kubernetes, Prometheus
- **Networks**: Ethereum Sepolia (11155111), Polygon Amoy (80002)

## 📚 Documentation
- [Architecture](./docs/architecture/) - System design and C4 models
- [API Specification](./docs/api/) - REST API documentation  
- [Analysis](./docs/analysis/) - Requirements and user stories
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

### 🔜 Phase 3: Go Relay Service (NEXT)
- Event listener for cross-chain transactions
- Signature generation and verification
- PostgreSQL for idempotency
- Docker containerization

## 📊 Smart Contracts Overview

### BridgeEthereum.sol
```solidity
function lockTokens(uint256 amount, address targetAddress, uint256 targetChainId)
```
- Locks tokens on Ethereum side
- Emits TokensLocked event for relayer
- Full input validation and security checks
  
### BridgePolygon.sol
```solidity
function releaseTokens(address user, uint256 amount, uint256 nonce, bytes memory signature)
```
- Releases tokens on Polygon side
- EIP-712 signature verification
- Nonce-based idempotency protection

🧪 Testing
```bash
cd contracts
forge test -vv
```
**Results**: 13/13 tests passing ✅

- Unit tests for all contracts
- Integration tests for cross-chain flow
- Security and edge case coverage

🚀 Quick Start
```bash
# Clone repository
git clone https://github.com/dmsus/crossChainBridge.git
cd crossChainBridge

# Test smart contracts
cd contracts
forge test -vv

# Explore documentation
open docs/architecture/c4-model.md
open docs/api/openapi.yaml
```
🔜 Next Steps
- Deploy to testnets (Issues #11, #12)
- Implement Go relay service (Phase 3)
- Setup monitoring and infrastructure

📄 License
MIT

