# Testing — Mechigram

## Objetivos

1. Proteger invariantes del dominio (cruces, backtracking)
2. Detectar regresiones en normalización ES
3. Smoke de outputs críticos (PDF)
4. Habilitar CI confiable

## Convenciones

| Regla | Detalle |
|-------|---------|
| Ubicación | `*_test.go` mismo paquete |
| Estilo | Table-driven con `t.Run` |
| Nombres | `Test<Func>_<scenario>` o columna `name` |
| Fixtures | Preferir construcción inline; `testdata/` solo si > 50 líneas |
| Random | **Nunca** `rand` global en tests; `rand.New(rand.NewSource(42))` |
| Parallel | `t.Parallel()` solo tests sin estado compartido |

## Cobertura objetivo (Iter. 2)

| Paquete | Meta |
|---------|------|
| `core` | ≥ 80 % |
| `dictionary` | ≥ 80 % |
| `render/pdf` | Smoke + 1 test layout |
| `internal/app` | 1 test integración con dict fixture |
| `cmd` | Opcional — thin wrapper |

## Qué testear / qué no

### Sí

- `CanPlaceWord` conflictos y cruces válidos
- `RemoveWord` preserva letras ajenas
- Parser: slots, min length 4, board vacío
- Solver: solución conocida, seed determinista
- `processWord`: tildes, Ñ, puntuación
- PDF: bytes `%PDF`, tamaño > 0

### No (Iter. 2)

- Snapshot visual pixel-perfect del PDF
- Performance benchmarks salvo issue explícito
- Tests que dependen de `input/words.txt` completo (usar dict mínimo)

## testdata/

```
core/testdata/
  mini_board.txt
  mini_words.txt
```

Solo fixtures estables y pequeños.

## Comandos

```bash
go test ./...
go test ./core/... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## CI

Todo PR debe pasar `go test ./...` y `go vet ./...`.

Ver [`playbooks/setup-ci/PLAYBOOK.md`](../playbooks/setup-ci/PLAYBOOK.md).
