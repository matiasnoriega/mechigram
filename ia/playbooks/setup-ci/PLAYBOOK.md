---
name: mechigram-setup-ci
description: Configura GitHub Actions y gates de calidad para Mechigram (go test, vet, build). Usar al crear .github/workflows, CI rojo, o alinear go.mod con versión Go del pipeline.
---

# Playbook: Setup CI

## Workflow mínimo

Crear `.github/workflows/ci.yml`:

```yaml
name: CI

on:
  push:
    branches: [main, master]
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: ['1.22.x', '1.23.x']
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
      - run: go vet ./...
      - run: go test ./...
      - run: go build -o /dev/null ./cmd/mechigram
```

**Nota:** Ajustar matrix cuando `go.mod` tenga versión canónica acordada.

## Pre-merge local

```bash
go vet ./...
go test ./...
go build ./cmd/mechigram
```

## golangci-lint (opcional)

```yaml
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: golangci/golangci-lint-action@v6
        with:
          version: latest
```

## CONTRIBUTING mínimo

En raíz, referenciar:

- Requiere CI verde
- `go test ./...` antes de push
- Ver `ia/standards/contributing.md`

## Verificación

- Push a branch → Actions verde
- Falla intencional en test → Actions rojo

## Referencias

- `ia/ROADMAP.md` — Épica 5
- `ia/roles/platform-engineer.md`
