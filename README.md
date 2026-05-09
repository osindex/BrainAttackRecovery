# BrainAttackRecovery

Self-hosted stroke rehabilitation training app for single-patient use.

## Layout

```text
.
├── docker-compose.yml        # Local stack: PostgreSQL + LinaPro + H5 nginx
├── infra/                    # Local LinaPro runtime config
├── h5/                       # Patient-facing H5 app (Vue 3 + Vant + PWA)
└── linapro/                  # Vendored LinaPro source fork with rehab plugins
```

## Current MVP

- Patient H5 with large-button flows for walking, fist-raise, eye-gaze, picture-card game, and history.
- Offline-first training records in IndexedDB, with sync queue to LinaPro when paired credentials are present.
- LinaPro source plugins:
  - `rehab-cards`: card categories, cards, and placeholder crawler jobs.
  - `rehab-records`: training records and daily aggregates.
- Local Docker stack uses PostgreSQL and loopback-only port bindings by default.
- Unified local entry: `http://127.0.0.1:18080`.
- Patient H5 is served from `/`; LinaPro API is proxied through `/api/*`.

## Local secrets

Copy `.env.example` to `.env` and replace every value before running Docker Compose.

```powershell
copy .env.example .env
docker compose up --build
```

## Security notes

- The H5 app no longer ships default admin credentials.
- Pair the H5 device in **本机配对** using a low-privilege LinaPro account created for that device.
- Do not use the built-in `admin/admin123` account for patient devices.
- `docker-compose.yml` binds service ports to `127.0.0.1` for local development.

## Verification

Validated locally:

- `go build ./apps/lina-core`
- `pnpm build` from `h5/`
- `docker compose config`

`docker compose build` may still fail in unauthenticated environments because Docker Hub can rate-limit base-image pulls with HTTP 429.
