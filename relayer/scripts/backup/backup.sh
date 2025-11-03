#!/bin/bash
# Database backup script
set -e

BACKUP_DIR="/backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/backup_$TIMESTAMP.sql"

echo "Creating database backup: $BACKUP_FILE"

pg_dump -h postgres -U bridge_user -d bridge_local -F c > $BACKUP_FILE

echo "Backup completed: $BACKUP_FILE"
