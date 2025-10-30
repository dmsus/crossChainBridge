# Deployment Guide for Cross-Chain Bridge

## Prerequisites

### 1. Get API Keys (FREE)
- **Infura**: https://infura.io/ (for RPC endpoints)
- **Etherscan**: https://etherscan.io/apis (for contract verification)
- **Polygonscan**: https://polygonscan.com/apis (for contract verification)

### 2. Get Testnet Tokens
- **Sepolia ETH**: https://sepoliafaucet.com/ (need Alchemy account)
- **Polygon Amoy MATIC**: https://faucet.polygon.technology/ (select Amoy network)

### 3. Setup Environment
```bash
# Copy example file and fill with your keys
cp .env.example .env

# Edit .env with your actual keys
nano .env
```
### Deployment Commands
### Deploy to Ethereum Sepolia
```bash
forge script script/DeployEthereum.s.sol:DeployEthereum --rpc-url sepolia --broadcast --verify -vvv
```
### Deploy to Polygon Amoy
```bash
forge script script/DeployPolygon.s.sol:DeployPolygon --rpc-url amoy --broadcast --verify -vvv
```
### Post-Deployment
Update documentation with deployed contract addresses

Test basic functionality on testnet

Prepare for Phase 3 (Go relayer implementation)

### Network Configuration
### Ethereum Sepolia
Chain ID: 11155111

RPC URL: https://sepolia.infura.io/v3/YOUR_INFURA_KEY

Explorer: https://sepolia.etherscan.io

### Polygon Amoy
Chain ID: 80002

RPC URL: https://polygon-amoy.infura.io/v3/YOUR_INFURA_KEY

Explorer: https://amoy.polygonscan.com
