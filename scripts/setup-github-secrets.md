# GitHub Secrets Setup Guide

## Required Secrets for CI/CD

### Staging Environment
- `STAGING_ETH_PRIVATE_KEY`: 0x95f1623c0508df464789e769da9ef7e16ae6d323d2a3e67ad8c1b4b94ada5f8c
- `STAGING_POLYGON_PRIVATE_KEY`: 0x95f1623c0508df464789e769da9ef7e16ae6d323d2a3e67ad8c1b4b94ada5f8c
- `SEPOLIA_RPC_URL`: https://sepolia.infura.io/v3/afa8f63b32a84f508f554b798b247453
- `AMOY_RPC_URL`: https://polygon-amoy.infura.io/v3/afa8f63b32a84f508f554b798b247453
- `ETHERSCAN_API_KEY`: EKBDVYF7WC6NTF8XMA35VSMHYJCPFRE5C8
- `POLYGONSCAN_API_KEY`: EKBDVYF7WC6NTF8XMA35VSMHYJCPFRE5C8

### Production Environment (Placeholder)
- `PRODUCTION_ETH_PRIVATE_KEY`: [MAINNET_PRIVATE_KEY]
- `PRODUCTION_POLYGON_PRIVATE_KEY`: [MAINNET_PRIVATE_KEY]
- `MAINNET_RPC_URL`: [MAINNET_RPC_URL]
- `POLYGON_MAINNET_RPC_URL`: [POLYGON_MAINNET_RPC_URL]

## Setup Instructions

1. Go to Repository Settings → Secrets and variables → Actions
2. Add each secret with the exact name and value
3. Configure environment protection rules for staging and production
