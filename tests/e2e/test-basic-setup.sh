#!/bin/bash
set -e

echo "🔧 Testing basic setup..."

# Проверяем правильный forge
export PATH="$HOME/.foundry/bin:$PATH"
echo "✅ Using forge: $(which forge)"
forge --version

# Проверяем anvil
echo "✅ Using anvil: $(which anvil)" 
anvil --version

# Запускаем тестовые сети в фоне
echo "⛓️ Starting test networks..."
pkill anvil || true
anvil --port 8545 --chain-id 31337 --silent &
anvil --port 8546 --chain-id 31338 --silent &

sleep 2

# Проверяем подключение
echo "🔌 Testing connections..."
curl -s -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' http://localhost:8545 | jq -r '.result' && echo "✅ Ethereum connected"
curl -s -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' http://localhost:8546 | jq -r '.result' && echo "✅ Polygon connected"

echo "🎉 Basic setup test passed!"
