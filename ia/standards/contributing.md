# Contributing — Mechigram

Gracias por contribuir. Este proyecto prioriza puzzles **imprimibles** y código **mantenible**.

## Antes de empezar

1. Lee [`ia/CONTEXT.md`](../CONTEXT.md)
2. Revisa [`ia/ROADMAP.md`](../ROADMAP.md) — ¿encaja en iteración actual?
3. Abre issue para cambios grandes

## Setup local

```bash
git clone https://github.com/matiasnoriega/mechigram.git
cd mechigram
go test ./...
go run .   # Iter. 1; migrará a go run ./cmd/mechigram
```

## Flujo de trabajo

1. Fork / branch desde `main`
2. Implementar con tests
3. `go test ./...` && `go vet ./...`
4. PR con descripción clara
5. CI verde obligatorio

## Scope de PR

- **Una preocupación por PR** cuando sea posible (ej. PDF separado de RNG)
- Actualizar docs/harness si cambia estructura o flags
- No mezclar formato unrelated (`gofmt` masivo + feature)

## Estándares

| Tema | Doc |
|------|-----|
| Arquitectura | [`architecture.md`](architecture.md) |
| Tests | [`testing.md`](testing.md) |
| Nombres | [`naming-conventions.md`](naming-conventions.md) |
| Go | [`go-style.md`](go-style.md) |

## Harness IA (`ia/`)

Si añades paquete, flag o rol:

- Actualizar `ia/CONTEXT.md` y `ia/MANIFEST.md`
- Considerar playbook en `ia/playbooks/`

## Definition of Done (contribución)

- [ ] Tests relevantes
- [ ] Sin romper `core` puro
- [ ] README o ia/ actualizado si aplica
- [ ] PR checklist [`workflows/pr-checklist.md`](../workflows/pr-checklist.md)

## Código de conducta

Interacción respetuosa. Proyecto personal con fines familiares y OSS — mantener tono profesional.

## Licencia

Al contribuir, aceptas licencia del repo (ver `LICENSE`).
