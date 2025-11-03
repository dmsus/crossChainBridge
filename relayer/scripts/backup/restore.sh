#!/bin/bash
# Database restore script
set -e

BACKUP_FILE=$1

if [ -z "$BACKUP_FILE" ]; then
    echo "Usage: $0 <backup_file>"
    exit 1
fi

echo "Restoring database from: $BACKUP_FILE"

pg_restore -h postgres -U bridge_user -d bridge_local -c $BACKUP_FILE

echo "Restore completed"
