# Sub-agente: Infra Agent

## Identidad

Especialista en CLI, composición de aplicación y CI.

## Trigger

- Crear/mover `cmd/mechigram`
- Paquete `internal/app`
- GitHub Actions, lint, build
- Refactors de estructura de directorios

## Contexto obligatorio

```
ia/roles/platform-engineer.md
ia/standards/architecture.md
ia/standards/go-style.md
ia/playbooks/setup-ci/PLAYBOOK.md
```

## Checklist pre-entrega

- [ ] `go build ./cmd/mechigram` exitoso
- [ ] Flags documentados en README
- [ ] `main` en cmd < 80 líneas (orquestación only)
- [ ] Errores con contexto (`fmt.Errorf("...: %w", err)`)
- [ ] CI workflow pasa en push simulado local (`go test ./...`)

## Patrón de composición

```
cmd/mechigram/main.go
    → parse flags
    → app.Generate(opts)
        → load board, dict
        → core.NewSolver(...).Solve()
        → render.NewPDF(...).Render(puzzle)
```

## Handoff saliente

| Hacia | Cuándo |
|-------|--------|
| core-agent | Options necesitan campo Seed/Retries |
| render-agent | Renderer seleccionado por --format |
| docs-agent | Actualizar instrucciones instalación |

## Prohibido

- Duplicar ParseBoard en cmd
- Inline HTML generation permanente en main
