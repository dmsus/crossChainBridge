# Production Environment

Deploys to main networks:
- Ethereum Mainnet
- Polygon Mainnet

## Required Secrets
- `PRODUCTION_ETH_PRIVATE_KEY`: Private key for Ethereum Mainnet deployment
- `PRODUCTION_POLYGON_PRIVATE_KEY`: Private key for Polygon Mainnet deployment
- `MAINNET_RPC_URL`: RPC endpoint for Ethereum Mainnet
- `POLYGON_MAINNET_RPC_URL`: RPC endpoint for Polygon Mainnet
- `ETHERSCAN_API_KEY`: API key for contract verification
- `POLYGONSCAN_API_KEY`: API key for contract verification

## Protection Rules
- Requires manual approval
- Only deploy from tagged releases
