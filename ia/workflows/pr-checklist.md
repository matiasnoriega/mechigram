# PR Checklist — Mechigram

Completar antes de abrir o aprobar un Pull Request.

## Contexto

- [ ] PR referencia issue o épica ROADMAP (ej. "ROADMAP 2.3")
- [ ] Scope acotado — no incluye trabajo de iteración futura

## Código

- [ ] `go test ./...` pasa localmente
- [ ] `go vet ./...` pasa
- [ ] `gofmt` aplicado
- [ ] Sin secrets (.env, keys) en diff
- [ ] `core/` no importa render/os/http

## Tests

- [ ] Comportamiento nuevo tiene test
- [ ] Tests usan seed fijo si hay random
- [ ] No tests flaky (sin sleep, sin orden no determinista)

## UX / Producto (si aplica render/CLI)

- [ ] Flags nuevos documentados
- [ ] Modo puzzle vs solución verificado
- [ ] Ñ / caracteres ES probados

## Docs

- [ ] README actualizado si cambia CLI o estructura
- [ ] `ia/CONTEXT.md` o MANIFEST si cambia arquitectura
- [ ] Comentarios exportados en Go actualizados

## Review

- [ ] Self-review del diff completo
- [ ] Sin `TODO` sin issue asociado (o justificado)
- [ ] Nombres siguen [`naming-conventions.md`](../standards/naming-conventions.md)

## CI

- [ ] GitHub Actions verde (cuando exista)

## Aprobación por rol (opcional)

| Rol | Requerido si PR toca… |
|-----|----------------------|
| domain-engineer | `core/`, `dictionary/` |
| render-engineer | `render/` |
| platform-engineer | `cmd/`, `.github/` |
| qa-engineer | solo tests — always nice |
| docs-agent | cambios estructurales |
