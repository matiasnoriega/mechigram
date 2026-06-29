# Arquitectura — Mechigram

## Capas y dependencias

```
                    ┌─────────────┐
                    │ cmd/mechigram│  CLI, flags, exit codes
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │ internal/app │  Orquestación, opts, reintentos
                    └──────┬──────┘
           ┌───────────────┼───────────────┐
           │               │               │
    ┌──────▼──────┐ ┌──────▼──────┐ ┌──────▼──────┐
    │ dictionary  │ │    core     │ │   render    │
    └─────────────┘ └─────────────┘ └─────────────┘
         (I/O            (puro)         (I/O
          carga)                        salida)
```

**Regla de oro:** flechas solo hacia abajo. `core` no importa `render`, `dictionary` no importa `core`.

## Paquetes

| Paquete | Responsabilidad | I/O |
|---------|-----------------|-----|
| `core` | Board, Parser, Solver | ❌ |
| `dictionary` | Carga y normalización | ✅ lectura archivo |
| `render` | PDF, HTML | ✅ escritura archivo |
| `internal/app` | Pipeline completo | ✅ compone I/O |
| `cmd/mechigram` | Entrypoint | ✅ stdin/stdout/stderr |

## Interfaces clave (Iter. 2)

```go
// render/renderer.go
type Puzzle struct {
    Board       *core.Board
    Words       []string
    ShowLetters bool
}

type Renderer interface {
    Render(p Puzzle, w io.Writer) error
}

// internal/app/app.go
type Options struct {
    BoardPath  string
    DictPath   string
    OutputPath string
    Format     string // "pdf" | "html"
    Seed       int64
    Solution   bool
    MaxRetries int
}

func Generate(opts Options) error
```

## Estado vs mutación

| Componente | Mutable durante solve |
|------------|----------------------|
| `Board` en solver | Sí (Place/Remove) |
| Diccionario `map[int][]string` | No — shuffle copia local |
| Plantilla en disco | No — re-parse o Clone por intento |

## Evolución futura (Iter. 3)

```
internal/batch/     # worker pool, merge PDFs
api/                # HTTP handlers → app.Generate
```

No crear estos paquetes hasta ROADMAP Iter. 3.

## Anti-patrones

- `core` importando `github.com/jung-kurt/gofpdf`
- Lógica de backtracking en `render/pdf`
- `main.go` monolítico en raíz (migrar a `cmd/`)

## Diagrama de secuencia — generación

```
CLI → app.Generate
  → ParseBoard
  → LoadDictionary
  → loop retries:
       → NewSolver(board, dict, rng).Solve()
  → build Puzzle
  → renderer.Render(puzzle, file)
```
