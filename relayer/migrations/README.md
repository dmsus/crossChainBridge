# Database Schema for Cross-Chain Bridge

## Connection Details
- **Database**: `bridge_local`
- **User**: `bridge_user` 
- **Password**: `bridge_password`
- **Host**: `localhost`
- **Port**: `5432`

## Tables Created
1. **`transactions`** - tracks cross-chain transfers with status management
2. **`processed_events`** - ensures idempotency by tracking processed blockchain events  
3. **`nonce_tracking`** - prevents replay attacks through nonce monitoring

## Apply Migrations
```bash
sudo -u postgres psql -d bridge_local -f 001_init_schema.sql
```
Schema Ready for Phase 3 Go Relay Service
