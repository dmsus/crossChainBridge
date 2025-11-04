#!/bin/bash
set -e

# Убедимся что используем правильный forge
export PATH="$HOME/.foundry/bin:$PATH"

echo "🚀 Deploying test contracts..."

# Ждем пока anvil запустится
echo "⏳ Waiting for Anvil nodes to start..."
sleep 3

# Деплой в Ethereum testnet
echo "📦 Deploying to Ethereum testnet..."
cd ../../contracts

PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
forge script script/DeployEthereum.s.sol:DeployEthereum \
  --rpc-url http://localhost:8545 \
  --broadcast \
  --skip-simulation

# Деплой в Polygon testnet  
echo "📦 Deploying to Polygon testnet..."
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
forge script script/DeployPolygon.s.sol:DeployPolygon \
  --rpc-url http://localhost:8546 \
  --broadcast \
  --skip-simulation

echo "✅ Test contracts deployed successfully!"
