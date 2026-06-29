# Rol: Platform Engineer

## Misión

Infraestructura del repo: CLI, estructura de paquetes, CI/CD, calidad de build. Hace que el dominio y el render sean **ejecutables y verificables** por cualquiera.

## Scope de archivos

```
cmd/**/*
internal/app/**/*
.github/**/*
go.mod
go.sum
Makefile                  # si se añade
.gitignore
```

## Responsabilidades

- Entrypoint `cmd/mechigram` con flags documentados
- Paquete `internal/app` que compone dictionary + core + render
- GitHub Actions: test, vet, build
- Alinear `go.mod` con versión Go real del CI
- No filtrar lógica de negocio en `cmd` — solo parsing de flags y errores UX

## Flags CLI objetivo (Iter. 2)

| Flag | Tipo | Default |
|------|------|---------|
| `--board` | path \| `random` | `input/board1.txt` |
| `--dict` | path | `input/words.txt` |
| `--output` | path | `output/cruzada.pdf` |
| `--format` | `pdf` \| `html` | `pdf` |
| `--seed` | int64 | `0` (aleatorio) |
| `--solution` | bool | `false` |

## Fuera de scope

- Heurísticas del solver
- Detalles de layout PDF (coordenadas mm)
- Contenido de plantillas de tablero

## Referencias

- [`standards/architecture.md`](../standards/architecture.md)
- [`playbooks/setup-ci/PLAYBOOK.md`](../playbooks/setup-ci/PLAYBOOK.md)
- [`ROADMAP.md`](../ROADMAP.md) — Épica 1 y 5
