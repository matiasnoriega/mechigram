# Rol: Domain Engineer

## Misión

Mantener la lógica de negocio de cruzadas: tablero, parser, solver y diccionario. Código puro, testeable, sin efectos secundarios ocultos.

## Scope de archivos

```
core/**/*.go
dictionary/**/*.go
input/board*.txt          # plantillas (contenido, no código)
input/boards/**/*
input/words.txt
```

## Responsabilidades

- `ParseBoard`: detección correcta de slots H/V (longitud ≥ 4)
- `CanPlaceWord`, `PlaceWord`, `RemoveWord`: invariantes de cruce
- `Solver`: backtracking, unicidad de palabras, RNG (Iter. 2)
- `dictionary`: normalización ES (Ñ, sin tildes, solo letras)
- Tests unitarios del dominio (ver [`playbooks/write-core-tests/`](../playbooks/write-core-tests/))

## Invariantes que nunca romper

1. `RemoveWord` solo vacía celdas donde `placed[i] == true`
2. Longitudes en **runes**, no bytes
3. Palabra usada una sola vez por puzzle
4. Solver no muta el diccionario compartido entre goroutines (futuro: copia o shuffle local)

## Fuera de scope

- PDF, HTML, fuentes (`render/`)
- Flags CLI, CI (`cmd/`, `.github/`)
- Documentación de harness salvo glosario en CONTEXT

## Patrones preferidos

```go
// RNG inyectado — no math/rand global en tests
func NewSolver(b *Board, dict map[int][]string, rng *rand.Rand) *Solver

// Board clonable para reintentos sin corromper plantilla
func (b *Board) Clone() *Board  // cuando se implemente
```

## Referencias

- [`standards/architecture.md`](../standards/architecture.md)
- [`standards/testing.md`](../standards/testing.md)
- [`playbooks/solve-puzzle/PLAYBOOK.md`](../playbooks/solve-puzzle/PLAYBOOK.md)
