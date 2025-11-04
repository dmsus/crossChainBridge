# 🚀 Quick Start: Local Development with Docker

This guide will help you set up the cross-chain bridge for local development using Docker and Anvil.

## Prerequisites

- Docker & Docker Compose
- Foundry (for smart contract development)
- Go 1.21+ (for relayer development)

## Step 1: Clone and Setup

```bash
git clone https://github.com/dmsus/crossChainBridge
cd crossChainBridge

# Install Foundry dependencies
cd contracts
forge install
```
## Step 2: Start Local Blockchain (Anvil)
```bash
# Start two Anvil instances simulating Ethereum and Polygon
anvil --chain-id 31337 --port 8545 &
anvil --chain-id 31338 --port 8546 &

# Deploy contracts to local networks
forge script script/DeployEthereum.s.sol --rpc-url http://localhost:8545 --broadcast
forge script script/DeployPolygon.s.sol --rpc-url http://localhost:8546 --broadcast
```
## Step 3: Start the Relayer
```bash
cd relayer

# Copy environment configuration
cp .env.local.yaml .env.local

# Start with Docker Compose
docker-compose up -d

# Or run locally
go run cmd/relayer/main.go --env local
```
## Step 4: Verify Setup
Check that all services are running:

```bash
# Check Docker containers
docker-compose ps

# Check health endpoints
curl http://localhost:8081/health
curl http://localhost:9090/metrics

# Check blockchain connections
curl -X POST http://localhost:8545 -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```
## Step 5: Test the Bridge
Use the provided test scripts to verify the bridge functionality:

```bash
# Run E2E tests
cd tests/e2e
./test-basic-setup.sh
```
## Troubleshooting
## Common Issues
1. Port conflicts: Change ports in docker-compose.yml

2. Database connection: Ensure PostgreSQL is running on port 5432

3. Blockchain connection: Verify Anvil instances are running

## Getting Help
- Check logs: docker-compose logs relayer

- View metrics: http://localhost:9090/metrics

- Health check: http://localhost:8081/health
