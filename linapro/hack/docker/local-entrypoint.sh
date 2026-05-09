#!/bin/sh
set -eu

CONFIG_DIR=/tmp/linapro-config
mkdir -p "$CONFIG_DIR"
sed \
  -e "s|\${POSTGRES_DSN}|${POSTGRES_DSN}|g" \
  -e "s|\${LINAPRO_JWT_SECRET}|${LINAPRO_JWT_SECRET}|g" \
  /app/config.yaml > "$CONFIG_DIR/config.yaml"
export GF_GCFG_PATH="$CONFIG_DIR"

./lina init --confirm=init
exec ./lina
