#!/usr/bin/env bash
set -euo pipefail

BACKUP_DIR=${BACKUP_DIR:-./backup}
DB_NAME=${DB_NAME:-linapro}
DB_USER=${DB_USER:-postgres}

mkdir -p "$BACKUP_DIR"
pg_dump -h 127.0.0.1 -p 5432 -U "$DB_USER" "$DB_NAME" > "$BACKUP_DIR/linapro-$(date +%F-%H%M%S).sql"
tar czf "$BACKUP_DIR/linapro-files-$(date +%F-%H%M%S).tar.gz" data/linapro-upload data/linapro-output

echo "Backup completed: $BACKUP_DIR"
