# Naming Conventions — Mechigram

## Idioma

| Ámbito | Idioma |
|--------|--------|
| Código (tipos, funciones, vars) | **Inglés** |
| Comentarios de dominio complejo | Español aceptable si aclara regla de negocio |
| Docs harness (`ia/`) | Español |
| README usuario | Español o bilingüe; CLI en inglés (flags) |

## Paquetes

- Minúsculas, una palabra: `core`, `render`, `dictionary`
- Sin `util`, `common`, `helpers`
- `internal/app` — no exportable fuera del module

## Archivos

| Tipo | Patrón | Ejemplo |
|------|--------|---------|
| Dominio | sustantivo.go | `board.go`, `solver.go` |
| Parser | verbo.go | `parser.go` |
| Tests | `*_test.go` | `board_test.go` |
| Plantillas | snake o numbered | `input/boards/classic_01.txt` |
| Output | usuario define | `output/cruzada.pdf` |

## Tipos

| Concepto | Nombre |
|----------|--------|
| Tablero | `Board` |
| Celda | `Cell` |
| Hueco | `Slot` |
| Dirección | `Direction`, `Horizontal`, `Vertical` |
| Resultado impresión | `Puzzle` |
| Opciones CLI | `Options` |

Evitar sinónimos: usar **Slot** (no Gap, Hole, Space).

## Funciones exportadas

- Verbos PascalCase: `ParseBoard`, `LoadDictionary`, `CanPlaceWord`
- Constructores: `NewSolver`, `NewPDFRenderer`
- Getters sin prefijo Get si idiomatic Go: `PlacedWords()` preferible a `GetPlacedWords()` en código nuevo (legacy existente puede quedar)

## Variables

- `i`, `j` loops cortos
- `rng` para `*rand.Rand`
- `runes` para `[]rune(word)`
- Evitar `dict` shadowing en scopes profundos — `dictionary` o `wordsByLen`

## Flags CLI

Kebab en CLI, camel en struct:

```
--board-path  →  BoardPath  (o --board según ROADMAP)
--show-solution → Solution
```

## Constantes

```go
const (
    MinSlotLength = 4
    DefaultMaxRetries = 10
)
```

## Git

| Item | Convención |
|------|------------|
| Branches | `feat/pdf-renderer`, `fix/solver-seed`, `chore/ci` |
| Commits | Imperativo, inglés: `add PDF renderer with puzzle mode` |
