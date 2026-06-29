---
name: mechigram-render-pdf
description: Genera o ajusta salida PDF imprimible para cruzadas Mechigram (A4, tipografía grande, modo puzzle/solución). Usar en render/pdf, layout impresión, fuentes TTF, o flag --format pdf.
---

# Playbook: Render PDF

## Precondiciones

- Contrato `Puzzle` definido (ver `ia/agents/render-agent.md`)
- Leer spec de layout en `ia/roles/render-engineer.md`

## Workflow

### 1. Scaffold paquete

```
render/
├── renderer.go      # type Renderer interface { Render(Puzzle) error }
├── pdf/
│   └── pdf.go
└── html/
    └── html.go
```

### 2. Dependencia

Default: `github.com/jung-kurt/gofpdf`

```bash
go get github.com/jung-kurt/gofpdf
```

Si Ñ falla con fuente built-in → embeber TTF en `assets/fonts/` y registrar en PDF.

### 3. Calcular geometría

```
pageW, pageH := 210.0, 297.0  // A4 mm
margin := 15.0
usable := pageW - 2*margin
cellSize := min(usable/board.Width, usable/board.Height, 14.0)
gridW := cellSize * board.Width
originX := (pageW - gridW) / 2
```

### 4. Dibujar grid

Por cada celda `(x,y)`:

- `#` → rectángulo relleno negro
- `' '` → rectángulo blanco con borde
- letra (modo solución) → centrada en celda

Modo puzzle: **no** dibujar letras aunque `Grid` las contenga.

### 5. Lista de palabras

Debajo del grid, fuente 14–16 pt, 2–3 columnas, palabras en `Puzzle.Words` (ya ordenadas).

### 6. Smoke test

```go
func TestPDFSmoke(t *testing.T) {
    // generar PDF en t.TempDir()
    // assert bytes.HasPrefix(data, []byte("%PDF"))
}
```

## Checklist impresión

- [ ] Imprimir prueba A4
- [ ] Legible a distancia
- [ ] Ñ en palabra de prueba

## Referencias

- `ia/ROADMAP.md` — Épica 2
