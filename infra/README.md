# Local Infrastructure

This directory stores local-only runtime configuration used by the root
`docker-compose.yml` stack.

## Services

- PostgreSQL 16 at `127.0.0.1:15432`
- LinaPro API at `http://127.0.0.1:8080`
- Patient H5 at `http://127.0.0.1:5173`

The LinaPro container reads `infra/linapro.config.yaml` via `GF_GCFG_PATH=/app`.

Before running the stack, copy the root `.env.example` to `.env` and replace all placeholder secrets.
