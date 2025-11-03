# CI/CD Workflows

## Overview

This project uses GitHub Actions for continuous integration and deployment.

## Workflows

### CI (Continuous Integration)
- **File**: `.github/workflows/ci.yml`
- **Trigger**: Push to `develop`, `feature/*` branches or PR to `main`
- **Purpose**: Fast validation and testing
- **Jobs**:
  - Quick checks (structure, YAML, Kubernetes)
  - Full test suite (for main branch PRs)

### CD (Continuous Deployment)  
- **File**: `.github/workflows/cd.yml`
- **Trigger**: Push to `main` or manual dispatch
- **Purpose**: Build, test, and deploy to environments
- **Jobs**:
  - Test and build Docker images
  - Security scanning
  - Code quality checks
  - Staging deployment (Sepolia/Amoy)
  - Production deployment (manual)

## Environments

### Staging
- **Networks**: Ethereum Sepolia, Polygon Amoy
- **Auto-deploy**: On push to main
- **Secrets Required**: Staging private keys, RPC URLs, API keys

### Production  
- **Networks**: Ethereum Mainnet, Polygon Mainnet
- **Deployment**: Manual approval required
- **Secrets Required**: Production private keys, mainnet RPC URLs

## Setup

1. Configure GitHub secrets using `scripts/setup-github-secrets.md`
2. Set up environments in repository settings
3. Configure branch protection rules
