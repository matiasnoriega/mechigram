# Integración con Cursor

Puente entre el harness `ia/` y las herramientas nativas de Cursor.

## Arquitectura dual

```
ia/                          .cursor/
├── CONTEXT.md               ├── rules/*.mdc      ← reglas concisas, alwaysApply/globs
├── standards/               └── skills/          ← opcional: copias de ia/playbooks para auto-discovery
├── playbooks/
├── agents/
└── roles/
```

**Fuente de verdad:** `ia/` (versionado, legible por humanos y cualquier agente).

**Cursor rules:** extractos accionables ≤ 50 líneas por regla, apuntando a `ia/`.

## Reglas recomendadas

Crear en `.cursor/rules/`:

| Archivo | alwaysApply | globs | Contenido |
|---------|-------------|-------|-----------|
| `mechigram-core.mdc` | false | `core/**/*.go` | Dominio puro, runes, PlaceWord recibo |
| `mechigram-render.mdc` | false | `render/**/*.go` | Layout impresión, Puzzle contract |
| `mechigram-global.mdc` | true | — | Leer ia/CONTEXT; no scope creep |

### Ejemplo `mechigram-global.mdc`

```markdown
---
description: Mechigram global context and constraints
alwaysApply: true
---

# Mechigram

Before non-trivial work, read `ia/CONTEXT.md` and relevant `ia/standards/`.

- Product: printable fill-in crosswords for 90+ user (33 Cruzadas style)
- `core/` must stay I/O-free
- Minimize diff scope; match existing Go conventions
- Full harness: `ia/README.md`
```

## Playbooks vs. Cursor Skills

Los **playbooks** viven en `ia/playbooks/` (vendor-neutral). Los **Cursor Skills** nativos viven en `.cursor/skills/` y se auto-descubren por frontmatter.

| Enfoque | Cuándo |
|---------|--------|
| **A — Referencia** (recomendada) | Playbooks en `ia/playbooks/`; invocar `@ia/playbooks/render-pdf/PLAYBOOK.md` |
| **B — Duplicar** | Copiar a `.cursor/skills/mechigram-render-pdf/SKILL.md` para auto-discovery en Cursor |

No duplicar sin política de sync — preferir A hasta que el equipo crezca.

## AGENTS.md en raíz

Cursor y otros tools detectan [`AGENTS.md`](../AGENTS.md) → apunta a `ia/README.md`.

## Task tool / subagentes

Al lanzar subagente en Cursor, incluir en el prompt:

```markdown
Full Repository Path: X:/WORKSPACE/mechigram
Agent definition: ia/agents/core-agent.md
Standards: ia/standards/architecture.md, ia/standards/testing.md
Task: [descripción con IDs ROADMAP]
```

## Mantenimiento

| Evento | Acción |
|--------|--------|
| Nuevo paquete | Actualizar `ia/standards/architecture.md` + rule glob |
| Nuevo playbook | Añadir a `ia/MANIFEST.md` |
| Iteración completada | ROADMAP + iteration checklist + CONTEXT estado |

## Verificación

1. Abrir archivo en `core/` → rule `mechigram-core` sugerida
2. Nueva conversación → rule global activa
3. Prompt "implement PDF" → agente debe consultar `ia/playbooks/render-pdf/`
