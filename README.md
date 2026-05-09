# BrainAttackRecovery

A self-hosted stroke rehabilitation H5 app for a single patient, with five
training modules and an admin-managed picture-card library.

## Layout

```
.
├── linapro/    # vendored fork of linaproai/linapro v0.1.0 (Go + GoFrame backend + Vben admin)
└── h5/        # patient-facing H5 (Vue 3 + Vant + Vite + PWA, to be scaffolded)
```

## Status

- repo bootstrapped
- linapro vendored (no upstream `.git`)
- h5 placeholder

## Training modules (planned)

1. Walking timer  — daily walking duration
2. Fist-raise counter  — sets × reps
3. Eye-gaze counter  — left-right repetitions
4. Picture-card naming  — image recognition mini-game
5. History dashboard  — cross-module charts

## Backend plugins to implement

- `rehab-cards`   — picture-card CRUD + categories + crawler (admin-side)
- `rehab-records` — patient training session records + aggregations

Both compile into the linapro host binary as source plugins.
