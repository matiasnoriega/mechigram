# Go Style — Mechigram

Convenciones Go además de [Effective Go](https://go.dev/doc/effective_go). Complementa [`naming-conventions.md`](naming-conventions.md).

## Errores

```go
// Envolver con contexto
if err != nil {
    return fmt.Errorf("load dictionary from %s: %w", path, err)
}

// Errores sentinela en app
var ErrNoSolution = errors.New("no solution found after max retries")
```

No `log.Fatal` fuera de `cmd/`.

## Receivers

- Structs con estado mutado (`Board`, `Solver`): pointer receiver
- Value receiver solo si struct pequeño e inmutable

## Imports

Grupos: stdlib → blank → third-party → local module.

```go
import (
    "fmt"
    "math/rand"

    "github.com/jung-kurt/gofpdf"

    "github.com/matiasnoriega/mechigram/core"
)
```

## Complejidad

- Funciones > 40 líneas: considerar extracción
- Backtracking recursivo en solver: aceptable; no convertir a iterativo sin motivo perf

## Unicode / Spanish

```go
runes := []rune(word)   // siempre para longitud y indexación
```

No indexar `word[i]` en lógica de juego.

## Comentarios

Exportados: sentence comment `// ParseBoard ...`

```go
// ParseBoard builds a Board from text rows where '#' marks black cells.
func ParseBoard(layout []string) *Board
```

No comentarios obvios (`// increment i`).

## Dependencias

- Preferir stdlib
- Nueva deps: justificar en PR; pure Go preferido
- Pin versions en `go.mod` via `go get`

## Formato

```bash
gofmt -w .
go vet ./...
```

golangci-lint cuando CI lo incluya.

## main / cmd

```go
func main() {
    if err := run(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

Lógica en `run()` testeable mínimamente.
