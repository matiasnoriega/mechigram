# Sub-agentes — Orquestación

Los sub-agentes son **unidades delegables** con contexto acotado. Un agente principal (orchestrator) los invoca según el tipo de tarea.

## Cuándo delegar vs. hacer inline

| Situación | Acción |
|-----------|--------|
| Tarea toca un solo paquete | Agente especializado directo |
| Tarea cruza core + render + cmd | Orchestrator coordina secuencia |
| Cambio < 20 líneas en un archivo | No delegar; aplicar rol inline |
| Épica completa de ROADMAP | Orchestrator + checklist |

## Reglas de handoff

1. **Contexto mínimo:** pasar rutas, IDs de épica ROADMAP, y estándares aplicables — no repetir CONTEXT entero
2. **Un owner por paquete** en un PR (evitar dos agentes editando `core/` a la vez)
3. **Orden de dependencias:** `core` → `render` → `cmd` → `docs`
4. **Sincronizar harness:** si un agente introduce estructura nueva, `docs-agent` actualiza `ia/CONTEXT.md` en el mismo ciclo

## Mapa agente → rol → archivos

| Agente | Rol equivalente | Paquetes |
|--------|-----------------|----------|
| [orchestrator](orchestrator.md) | tech-lead | todos (coordina) |
| [core-agent](core-agent.md) | domain-engineer | `core/`, `dictionary/` |
| [render-agent](render-agent.md) | render-engineer | `render/` |
| [infra-agent](infra-agent.md) | platform-engineer | `cmd/`, `internal/`, `.github/` |
| [docs-agent](docs-agent.md) | tech-lead (docs) | `README`, `ia/`, `CONTRIBUTING` |

## Secuencia típica — Iteración 2 Épica 1→2

```
1. infra-agent   → cmd/ + internal/app + interface Renderer (stub)
2. core-agent    → sin cambios o RNG prep
3. render-agent  → render/pdf implementación
4. qa-engineer   → tests smoke
5. docs-agent    → README flags
6. orchestrator  → verificar DoD ROADMAP
```

## Formato de prompt de delegación

```markdown
## Delegación: [core-agent]

**Épica:** ROADMAP 3.1–3.2
**Objetivo:** Inyectar RNG y shuffle en Solver
**Archivos:** core/solver.go, core/solver_test.go
**Leer:** ia/standards/testing.md, ia/playbooks/solve-puzzle/PLAYBOOK.md
**No tocar:** render/, cmd/
**DoD:** Test con seed fijo pasa; dos seeds → distintas soluciones en fixture
```

## Anti-patrones

- ❌ render-agent modificando `PlaceWord`
- ❌ infra-agent embebiendo HTML en `main.go` permanentemente
- ❌ Paralelizar edits en `internal/app` y `render/` sin contrato `Puzzle` acordado
- ❌ Delegar sin citar estándares → drift de estilo
