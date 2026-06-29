# Checklist — Iteración 2

Marcar conforme se complete. Sincronizar con [`ROADMAP.md`](../ROADMAP.md).

## Épica 1 — Arquitectura y CLI

- [ ] `cmd/mechigram/main.go` creado
- [ ] `internal/app` con `Generate(opts)`
- [ ] Interface `Renderer`
- [ ] Flags: `--board`, `--dict`, `--output`, `--seed`, `--format`
- [ ] Tipo `Puzzle` con palabras alfabéticas
- [ ] `main.go` raíz eliminado o deprecated

## Épica 2 — PDF

- [ ] Paquete `render/pdf`
- [ ] Paquete `render/html` (extraído)
- [ ] Modo puzzle (sin letras)
- [ ] Flag `--solution`
- [ ] Fuente con Ñ validada
- [ ] Prueba impresión real documentada

## Épica 3 — Variabilidad

- [ ] RNG en Solver
- [ ] Shuffle candidatos
- [ ] Reintentos en app
- [ ] `input/boards/` con ≥ 3 plantillas
- [ ] `--board random` o equivalente
- [ ] Verificado: 2 runs → palabras distintas

## Épica 4 — Tests

- [ ] `core/board_test.go`
- [ ] `core/parser_test.go`
- [ ] `core/solver_test.go`
- [ ] `dictionary/dictionary_test.go`
- [ ] `render/pdf` smoke test
- [ ] Cobertura core ≥ 80%

## Épica 5 — OSS

- [ ] `.github/workflows/ci.yml`
- [ ] README ampliado
- [ ] `CONTRIBUTING.md` en raíz (puede enlazar ia/)
- [ ] Issue templates (opcional)

## Harness

- [ ] `ia/CONTEXT.md` refleja estructura final
- [ ] `ia/ROADMAP.md` ítems marcados [x]
- [ ] `.cursor/rules` sincronizadas (ver CURSOR_INTEGRATION)

## Sign-off Iteración 2

- [ ] Tech Lead: DoD ROADMAP cumplido
- [ ] QA: CI verde
- [ ] Render: PDF aprobado para impresión
- [ ] Docs: README usable por contribuidor externo
