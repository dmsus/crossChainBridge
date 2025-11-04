# 💾 Backup & Recovery Procedures

Comprehensive backup and disaster recovery procedures for the cross-chain bridge.

## 📋 Backup Strategy

### Database Backups

**Automated Daily Backups**:
```bash
# Full database backup (runs daily at 2 AM)
pg_dump -h localhost -U bridge_user -d bridge_local -F c -f /backups/bridge_$(date +%Y%m%d).dump

# Backup retention: 7 daily, 4 weekly, 12 monthly
```
## Manual Backup:

```bash
cd relayer/scripts/backup
./backup.sh
```
## Configuration Backups
Backup these critical files:

- relayer/.env.*.yaml - Environment configurations

- contracts/broadcast/ - Deployment records

- k8s/ - Kubernetes manifests

- migrations/ - Database schema

## Smart Contract Backups
- Source code in contracts/src/

- Deployment addresses in contracts/broadcast/

- ABI files in contracts/out/

## 🔄 Recovery Procedures
## Database Recovery
## Full Database Restore:

```bash
# Stop the application first
docker-compose down

# Restore from backup
pg_restore -h localhost -U bridge_user -d bridge_local -c /backups/bridge_20241104.dump

# Restart application
docker-compose up -d
```
## Using Recovery Script:

```bash
cd relayer/scripts/backup
./restore.sh /backups/bridge_20241104.dump
```
## Partial Recovery Scenarios
## Recover Failed Transactions
```sql
-- Mark transactions for retry
UPDATE transactions 
SET status = 'pending', error_message = NULL, retry_count = 0
WHERE status = 'failed' AND created_at > NOW() - INTERVAL '1 hour';
```
## Reset Processing Queue
```sql
-- Reset stuck processing transactions
UPDATE transactions 
SET status = 'pending'
WHERE status = 'processing' AND updated_at < NOW() - INTERVAL '10 minutes';
```
## 🚨 Disaster Recovery
## Complete System Failure
## 1. Restore Infrastructure:

```bash
# Redeploy from Kubernetes manifests
kubectl apply -k k8s/production
```
## 2. Restore Database:

```bash
# Restore latest backup
./restore.sh /backups/bridge_latest.dump
```
## 3. Verify Smart Contracts:

```bash
# Verify contracts are still deployed
cast call --rpc-url $ETH_RPC_URL $BRIDGE_ETH "owner()"
```
## 4. Restart Services:

```bash
docker-compose up -d
# OR
kubectl rollout restart deployment/relayer
```
## Data Corruption Recovery
## Transaction Recovery Process:

1. Identify corrupted transactions

2. Export good transactions from backup

3. Import into clean database

4. Reconcile any gaps

## 🛡️ Security Backups
## Private Key Backup
**Critical**: Backup relayer private keys securely:

- Use encrypted storage

- Multiple secure locations

- Access restricted to authorized personnel

## API Key Rotation
Rotate API keys quarterly:

1. Generate new keys

2. Update configurations

3. Deploy with zero downtime

4. Invalidate old keys after 7 days

## 📊 Recovery Testing
## Quarterly DR Drill
**Test Scenario**:

1. Simulate database failure

2. Execute recovery procedures

3. Measure recovery time

4. Validate data integrity

## Success Criteria:

- Recovery Time Objective (RTO): < 4 hours

- Recovery Point Objective (RPO): < 15 minutes

- Data integrity: 100% transaction consistency

## Backup Validation
**Monthly Verification**:

```bash
# Verify backup integrity
pg_restore -l /backups/bridge_latest.dump | head -10

# Test restore on staging
./restore.sh --test /backups/bridge_latest.dump
```
## 📝 Recovery Checklists
## Immediate Actions (First 15 minutes)
- Assess impact scope

- Notify stakeholders

- Begin recovery procedures

- Document all actions

## Within 1 Hour
- Restore database from backup

- Verify service functionality

- Monitor system health

- Update status page

## Post-Recovery
- Root cause analysis

- Update procedures if needed

- Schedule follow-up review
