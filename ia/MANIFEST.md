# Manifest — Artefactos del Harness

Inventario versionado de todo el contexto para agentes. Actualizar cuando se añadan roles, playbooks o estándares.

## Documentos raíz (`ia/`)

| Archivo | Propósito | Audiencia |
|---------|-----------|-----------|
| `README.md` | Punto de entrada y mapa | Todos |
| `MANIFEST.md` | Este índice | Maintainers |
| `CONTEXT.md` | Visión, dominio, glosario, estado del código | Todos |
| `ROADMAP.md` | Plan por iteraciones | Tech Lead, Orchestrator |
| `CURSOR_INTEGRATION.md` | Reglas Cursor derivadas del harness | Maintainers |

## Roles (`ia/roles/`)

| Archivo | Responsabilidad principal |
|---------|---------------------------|
| `tech-lead.md` | Priorización, arquitectura, scope de iteraciones |
| `domain-engineer.md` | Board, Parser, Solver, diccionario |
| `render-engineer.md` | PDF/HTML, layout impresión, tipografía |
| `platform-engineer.md` | CLI, estructura de paquetes, CI/CD |
| `qa-engineer.md` | Tests, cobertura, calidad de entrega |

## Sub-agentes (`ia/agents/`)

| Archivo | Delegación típica |
|---------|-------------------|
| `README.md` | Orquestación, handoffs, reglas de delegación |
| `orchestrator.md` | Coordinador multi-épica |
| `core-agent.md` | Cambios en `core/` y `dictionary/` |
| `render-agent.md` | Cambios en `render/` |
| `infra-agent.md` | `cmd/`, CI, estructura repo |
| `docs-agent.md` | README, CONTRIBUTING, docs de harness |

## Playbooks (`ia/playbooks/`)

Workflows vendor-neutral (no confundir con skills nativos de Cursor/Claude). Ver [`README.md`](README.md#playbooks-vs-skills-nativos).

| Playbook | Trigger |
|----------|---------|
| `solve-puzzle/` | Implementar o modificar lógica de resolución |
| `render-pdf/` | Crear o ajustar salida PDF imprimible |
| `add-board-template/` | Nueva plantilla de tablero |
| `write-core-tests/` | Tests unitarios del dominio |
| `setup-ci/` | GitHub Actions y gates de calidad |

## Estándares (`ia/standards/`)

| Archivo | Contenido |
|---------|-----------|
| `architecture.md` | Capas, dependencias, interfaces |
| `testing.md` | Estrategia, cobertura, convenciones de tests |
| `naming-conventions.md` | Nombres de paquetes, tipos, archivos |
| `go-style.md` | Estilo Go específico del proyecto |
| `contributing.md` | Flujo git, PRs, definición de done |

## Workflows (`ia/workflows/`)

| Archivo | Uso |
|---------|-----|
| `iteration-2-checklist.md` | Seguimiento de la iteración actual |
| `pr-checklist.md` | Gate antes de abrir/revisar PR |

## Estado de sincronización

| Artefacto | Última alineación con código |
|-----------|------------------------------|
| `CONTEXT.md` | Iteración 1 (MVP) |
| `ROADMAP.md` | Plan Iteración 2 definido |
| Estructura objetivo `cmd/`, `render/` | **Pendiente** — descrita en ROADMAP, no implementada |

Cuando el código diverja del harness, **actualizar el harness en el mismo PR** que introduce el cambio estructural.
