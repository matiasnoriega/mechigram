---
name: mechigram-write-core-tests
description: Escribe tests unitarios table-driven para core/ y dictionary/ en Mechigram. Usar al añadir *_test.go, cubrir board/parser/solver, o subir cobertura del dominio.
---

# Playbook: Write Core Tests

## Estándar

Seguir `ia/standards/testing.md`.

## Plantilla table-driven

```go
func TestCanPlaceWord(t *testing.T) {
    tests := []struct {
        name  string
        board *Board
        word  string
        slot  Slot
        want  bool
    }{
        {name: "empty slot fits", ...},
        {name: "crossing letter matches", ...},
        {name: "crossing letter conflicts", ...},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.board.CanPlaceWord(tt.word, tt.slot)
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Fixtures mínimos

### Board 5×5 con cruce

Construir grid manualmente en test — no depender de archivos externos para unit tests rápidos.

### Parser

Usar layout string multilínea:

```go
layout := []string{
    ".....",
    ".###.",
    ".....",
}
```

### Solver

Dict mínimo controlado:

```go
dict := map[int][]string{
    4: {"CASA", "MESA", "PASA"},
    5: {"SALON", "CASAS"},
}
rng := rand.New(rand.NewSource(42))
```

## Cobertura prioritaria

1. `RemoveWord` no borra letra de palabra cruzada
2. Parser omite length < 4
3. `processWord` preserva Ñ
4. Solver seed fijo

## Comandos

```bash
go test ./core/... ./dictionary/... -cover
go test ./core/... -run TestName -v
```

## Referencias

- `ia/roles/qa-engineer.md`
