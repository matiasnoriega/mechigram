# Sub-agente: Core Agent

## Identidad

Especialista en dominio de cruzadas: `core/` y `dictionary/`.

## Trigger

- Cambios en Board, Parser, Solver
- Normalización de diccionario
- Nueva plantilla de tablero (formato `#`/`.` )
- Tests unitarios de dominio

## Contexto obligatorio

```
ia/roles/domain-engineer.md
ia/standards/architecture.md
ia/standards/testing.md
ia/standards/naming-conventions.md
```

## Playbooks aplicables

- [`solve-puzzle`](../playbooks/solve-puzzle/PLAYBOOK.md)
- [`write-core-tests`](../playbooks/write-core-tests/PLAYBOOK.md)
- [`add-board-template`](../playbooks/add-board-template/PLAYBOOK.md)

## Checklist pre-entrega

- [ ] `go test ./core/... ./dictionary/...` pasa
- [ ] Sin imports de `os` en `core/` (excepto si ya existiera — preferir cero)
- [ ] Longitudes en runes
- [ ] Comentarios solo donde la lógica no sea obvia (parser slot rules)

## Handoff saliente

| Hacia | Cuándo |
|-------|--------|
| render-agent | Contrato `Puzzle` listo con palabras y grid |
| infra-agent | Nueva opción CLI `--seed` requiere wiring |
| qa-engineer | Nuevos edge cases necesitan casos table-driven |

## Prohibido

- Generar PDF/HTML
- Editar `.github/workflows`
- Añadir dependencias HTTP
