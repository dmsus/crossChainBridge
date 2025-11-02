# 🔧 Configuration Management

## Overview
The Cross-Chain Bridge uses environment-based configuration with validation and support for multiple deployment environments.

## Environments
- **dev**: Local development with Anvil
- **staging**: Test networks (Sepolia/Amoy)  
- **production**: Mainnet deployment

## Quick Start

```bash
# Development
APP_ENV=dev go run cmd/relayer/main.go

# Staging (default)
go run cmd/relayer/main.go

# Staging with env var
APP_ENV=staging go run cmd/relayer/main.go
```
## Configuration Structure
# Core Configuration
```yaml
environment: "staging"  # dev, staging, production
```
## Database Configuration
```yaml
database:
  host: "localhost"           # PostgreSQL host
  port: "5432"               # PostgreSQL port
  user: "bridge_user"        # Database user
  password: "bridge_password" # Database password
  name: "bridge_local"       # Database name
  ssl_mode: "disable"        # SSL mode
  max_conns: 25              # Maximum connections
```
## Ethereum Configuration
```yaml
ethereum:
  rpc_url: "https://..."     # Ethereum RPC endpoint
  ws_url: "wss://..."        # Ethereum WebSocket endpoint  
  bridge_addr: "0x..."       # Bridge contract address
  token_addr: "0x..."        # Token contract address
  chain_id: 11155111         # Chain ID
  private_key: "0x..."       # Relayer private key
  block_confirmations: 3     # Block confirmations required
```
## Polygon Configuration
```yaml
polygon:
  rpc_url: "https://..."     # Polygon RPC endpoint
  ws_url: "wss://..."        # Polygon WebSocket endpoint
  bridge_addr: "0x..."       # Bridge contract address
  token_addr: "0x..."        # Token contract address  
  chain_id: 80002            # Chain ID
  private_key: "0x..."       # Relayer private key
  gas_limit: 500000          # Gas limit for transactions
``` 
## Processor Configuration
```yaml
processor:
  max_retries: 5             # Maximum retry attempts
  retry_delay: "10s"         # Delay between retries
  batch_size: 50             # Batch processing size
  queue_size: 500            # Event queue size
  health_check_period: "60s" # Health check interval
```
## Monitoring Configuration
```yaml
monitoring:
  enabled: true              # Enable monitoring
  metrics_port: 9090         # Prometheus metrics port
  log_level: "info"          # Log level (debug, info, warn, error)
  log_format: "json"         # Log format (text, json)
```
## API Configuration
```yaml
api:
  enabled: true              # Enable API server
  port: 8080                 # API port
  swagger_enabled: false     # Enable Swagger UI
  rate_limit: 1000           # Requests per minute
```
## Backward Compatibility
The configuration system maintains backward compatibility with the old ServiceConfig structure:

```yaml
# Old structure (still supported)
service:
  log_level: "info"
  metrics_port: "9090"
  api_port: "8080"

# New structure (recommended)
monitoring:
  log_level: "info"
  metrics_port: 9090
api:
  port: 8080
```
## Environment Variables
All configuration values can be overridden by environment variables:

```bash
# Application
export APP_ENV="staging"

# Database
export DATABASE_HOST="localhost"
export DATABASE_USER="bridge_user"
export DATABASE_PASSWORD="bridge_password"

# Ethereum
export ETHEREUM_RPC_URL="https://sepolia.infura.io/v3/your-key"
export ETHEREUM_PRIVATE_KEY="0x..."

# Polygon  
export POLYGON_RPC_URL="https://polygon-amoy.infura.io/v3/your-key"
export POLYGON_PRIVATE_KEY="0x..."
```
## Validation
The configuration system validates:

✅ Required fields presence

✅ Ethereum address format (0x + 40 hex chars)

✅ Private key format (64 hex chars)

✅ URL format for RPC endpoints

✅ Positive integers for numeric values

## Security Best Practices
1. Never commit private keys to version control

2. Use environment variables for production secrets

3. Enable SSL for database connections in production

4. Use different keys for different environments

5. Regularly rotate keys and credentials

## Example: Development Setup
1. Start PostgreSQL and Anvil

2. Create .env.dev.yaml with development settings

3. Run: APP_ENV=dev go run cmd/relayer/main.go

## Example: Production Setup
1. Set environment variables for all sensitive data

2. Create .env.production.yaml with variable references

3. Run: APP_ENV=production ./relayer
