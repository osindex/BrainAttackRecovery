#!/usr/bin/env bash
set -euo pipefail

if [ $# -lt 1 ]; then
  echo "Usage: $0 <backup.sql> [files.tar.gz]"
  exit 1
fi

DB_NAME=${DB_NAME:-linapro}
DB_USER=${DB_USER:-postgres}

psql -h 127.0.0.1 -p 5432 -U "$DB_USER" "$DB_NAME" < "$1"

if [ $# -ge 2 ]; then
  tar xzf "$2"
fi

echo "Restore completed"
