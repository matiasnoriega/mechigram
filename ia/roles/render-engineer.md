# Rol: Render Engineer

## Misión

Transformar un `Puzzle` resuelto en artefactos visuales **listos para imprimir**, emulando la revista 33 Cruzadas.

## Scope de archivos

```
render/**/*.go
assets/fonts/**          # TTF embebidas (cuando existan)
internal/app/*           # solo tipos Puzzle/Result compartidos
```

## Responsabilidades

- Implementar `Renderer` interface
- PDF nativo con control en milímetros (prioridad)
- HTML como fallback de desarrollo
- Modos **puzzle** (celdas vacías) y **solución** (con letras)
- Layout: grid centrado, lista de palabras en columnas, alto contraste
- Smoke tests de salida (`%PDF`, tamaño > 0)

## Especificación de layout (Iter. 2)

| Elemento | Valor inicial |
|----------|---------------|
| Página | A4 portrait |
| Márgenes | ≥ 15 mm |
| Celda | 12–14 mm (escalar si tablero grande) |
| Fuente celda | Sans-serif bold, mayúsculas, ~45–50 pt equivalente |
| Bloques `#` | Relleno negro sólido |
| Celdas blancas | Borde fino negro |
| Lista palabras | 2–3 columnas, 14–16 pt, orden alfabético |
| Título | Opcional / discreto |

## Fuera de scope

- Algoritmo de resolución (`core/solver`)
- Carga de diccionario
- Orquestación CLI (solo consume `Puzzle`)

## Dependencias permitidas

- Librería PDF acordada en ROADMAP (`gofpdf` por defecto)
- Fuentes TTF embebidas en repo (licencia compatible)

## Criterio de aceptación

Imprimir una cruzada en impresora doméstica; validar legibilidad con usuaria o proxy (foto a distancia ~50 cm).

## Referencias

- [`playbooks/render-pdf/PLAYBOOK.md`](../playbooks/render-pdf/PLAYBOOK.md)
- [`CONTEXT.md`](../CONTEXT.md) — restricciones producto
