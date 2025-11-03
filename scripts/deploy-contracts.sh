#!/bin/bash
set -e

echo "🚀 Starting contract deployment..."

cd contracts

echo "📦 Deploying to Ethereum Sepolia..."
PRIVATE_KEY="$1" forge script script/DeployEthereum.s.sol:DeployEthereum --rpc-url "$2" --broadcast --verify -vvvv

echo "📦 Deploying to Polygon Amoy..."  
PRIVATE_KEY="$3" forge script script/DeployPolygon.s.sol:DeployPolygon --rpc-url "$4" --broadcast --verify -vvvv

echo "🎉 ALL CONTRACTS DEPLOYED SUCCESSFULLY!"
