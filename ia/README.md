# Harness Engineering — Mechigram

Directorio de contexto estructurado para agentes de IA y colaboradores. Objetivo: que cualquier sesión (humana o automatizada) produzca código alineado con la visión del producto y los estándares del repo **sin depender de memoria conversacional**.

## Cómo usar este harness

### Para un agente de IA

1. **Siempre** leer [`CONTEXT.md`](CONTEXT.md) al iniciar una tarea no trivial.
2. Identificar el **rol** más cercano en [`roles/`](roles/) y actuar según sus responsabilidades y límites.
3. Consultar **estándares** en [`standards/`](standards/) antes de escribir código.
4. Si la tarea encaja con un **playbook** existente, seguir su `PLAYBOOK.md` paso a paso.
5. Para trabajo paralelo o delegación, usar las definiciones en [`agents/`](agents/).
6. Verificar entrega contra [`workflows/pr-checklist.md`](workflows/pr-checklist.md).

### Para un colaborador humano

- Empieza por [`CONTEXT.md`](CONTEXT.md) y [`ROADMAP.md`](ROADMAP.md).
- Antes de abrir un PR, revisa [`standards/contributing.md`](standards/contributing.md).
- Si añades una capacidad repetible, considera un nuevo playbook en [`playbooks/`](playbooks/).

## Playbooks vs. skills nativos

En `ia/playbooks/` viven **playbooks** (runbooks / workflows para agentes): instrucciones paso a paso, vendor-neutral, legibles por cualquier IA si se referencian en el prompt (p. ej. `@ia/playbooks/render-pdf/PLAYBOOK.md`).

**No** son skills nativos de un producto concreto:

| Producto | Skill nativo | Dónde vive |
|----------|--------------|------------|
| **Cursor** | Auto-discovery por descripción | `.cursor/skills/<nombre>/SKILL.md` |
| **Claude (Anthropic)** | Módulo empaquetado con scripts opcionales | Formato Skills de Anthropic |

Los playbooks de Mechigram son la **fuente de verdad** del harness. Cursor puede consumirlos vía reglas en [`.cursor/rules/`](../.cursor/rules/) o duplicando selectivamente en `.cursor/skills/` — ver [`CURSOR_INTEGRATION.md`](CURSOR_INTEGRATION.md).

## Mapa del directorio

```
ia/
├── README.md                 ← estás aquí
├── MANIFEST.md               ← índice de todos los artefactos
├── CONTEXT.md                ← visión, dominio, estado actual
├── ROADMAP.md                ← iteraciones y épicas
├── CURSOR_INTEGRATION.md     ← puente con .cursor/rules
├── roles/                    ← personas / responsabilidades
├── agents/                   ← sub-agentes delegables
├── playbooks/                ← workflows ejecutables (vendor-neutral)
├── standards/                ← normas de calidad
└── workflows/                ← checklists por fase
```

## Principios del harness

| Principio | Significado |
|-----------|-------------|
| **Contexto explícito** | Lo que no está escrito no existe para el agente |
| **Límites claros** | Cada rol/agente tiene un *scope* y un *out of scope* |
| **Progressive disclosure** | Playbooks concisos; detalle en standards y CONTEXT |
| **Producto primero** | La usuaria final (90+) manda sobre el portfolio |
| **Dominio puro** | `core/` no conoce I/O ni render |

## Índice rápido por tipo de tarea

| Quiero… | Ir a… |
|---------|-------|
| Entender el proyecto | [`CONTEXT.md`](CONTEXT.md) |
| Planificar iteración | [`ROADMAP.md`](ROADMAP.md) |
| Tocar solver / board / parser | [`roles/domain-engineer.md`](roles/domain-engineer.md) + [`agents/core-agent.md`](agents/core-agent.md) |
| Generar PDF / HTML | [`roles/render-engineer.md`](roles/render-engineer.md) + [`playbooks/render-pdf/PLAYBOOK.md`](playbooks/render-pdf/PLAYBOOK.md) |
| Tests y CI | [`roles/qa-engineer.md`](roles/qa-engineer.md) + [`playbooks/setup-ci/PLAYBOOK.md`](playbooks/setup-ci/PLAYBOOK.md) |
| Reestructurar paquetes / CLI | [`roles/platform-engineer.md`](roles/platform-engineer.md) + [`agents/infra-agent.md`](agents/infra-agent.md) |
| Coordinar una épica completa | [`roles/tech-lead.md`](roles/tech-lead.md) + [`agents/orchestrator.md`](agents/orchestrator.md) |
