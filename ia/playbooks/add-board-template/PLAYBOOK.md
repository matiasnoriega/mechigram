---
name: mechigram-add-board-template
description: Añade y valida plantillas de tablero de cruzadas Mechigram (archivos texto con # y espacios). Usar al crear input/boards/, nuevas formas de grid, o flag --board random.
---

# Playbook: Add Board Template

## Formato de plantilla

```
....#
.##.#
.....
.##.#
....#
```

| Carácter | Significado |
|----------|-------------|
| `#` | Bloque negro |
| `.` o espacio | Celda blanca |
| Filas | Misma longitud (rectangular) |

Guardar en `input/boards/<nombre>.txt`.

## Workflow

### 1. Diseñar layout

- Palabras mínimo **4 letras** (parser ignora slots más cortos)
- Verificar que existen slots H y V suficientes
- Evitar tableros demasiado grandes para A4 (> ~15×15 revisar escala PDF)

### 2. Validar parser

```go
board := core.ParseBoard(lines)
require.NotNil(t, board)
require.NotEmpty(t, board.Slots)
```

Imprimir slots en test de desarrollo:

```go
for _, s := range board.Slots {
    t.Logf("slot %d: len=%d dir=%v at (%d,%d)", s.ID, s.Length, s.Direction, s.StartX, s.StartY)
}
```

### 3. Validar solvabilidad

Correr generador con diccionario real:

```bash
go run ./cmd/mechigram --board input/boards/nueva.txt --dict input/words.txt
```

Si falla repetidamente → ajustar forma o ampliar diccionario.

### 4. Registrar

- Añadir archivo a `input/boards/`
- Actualizar README listando plantillas
- ROADMAP 3.7: mínimo 3 plantillas

## Cargador (cuando exista)

```go
func LoadRandomBoard(dir string, rng *rand.Rand) (*Board, error)
```

## Referencias

- `core/parser.go`
- `ia/CONTEXT.md` — glosario Slot
