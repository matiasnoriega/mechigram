# Contexto del Proyecto — Mechigram

## Visión

**Mechigram** es un generador dinámico de juegos de **Cruzadas** (Fill-in puzzles) en Go.

### Objetivo de producto

La usuaria final es una persona mayor de **90 años** que resuelve cruzadas impresas. El producto debe emular la experiencia de la revista **"33 Cruzadas"** (Edigrama):

- Salida optimizada para **impresión en papel**
- Tipografía **grande**, **alto contraste**, disposición clara
- Lista de palabras legible (columnas, sin depender del navegador)

### Objetivo técnico

Repositorio **open source** como portfolio de un Full Stack Engineer senior:

- Arquitectura limpia y mantenible
- Alto rendimiento donde importa (solver, generación en lote futura)
- Preparado para contribuciones externas
- `core` consumible por CLI, API o serverless sin reescritura

## Glosario de dominio

| Término | Definición |
|---------|------------|
| **Board** | Tablero completo: grid + slots detectados |
| **Cell** | Celda: `' '` vacía, `'#'` bloque negro, o letra colocada |
| **Slot** | Hueco donde encaja **exactamente una** palabra (H o V, longitud ≥ 4) |
| **Plantilla** | Archivo texto con `#` = negro, `.` o espacio = blanco |
| **Diccionario** | Palabras agrupadas por longitud (en runes, no bytes) |
| **Solver** | Backtracking que llena slots sin repetir palabras |
| **Puzzle** | Resultado: tablero resuelto + lista de palabras para el jugador |
| **Modo puzzle** | PDF/HTML **sin letras** en celdas (para resolver) |
| **Modo solución** | PDF/HTML **con letras** (validación / uso personal) |

## Estado actual (Iteración 1 — implementado)

```
mechigram/
├── main.go              # Entry monolítico + renderHTML inline
├── core/
│   ├── board.go         # Board, Cell, Slot, PlaceWord, RemoveWord
│   ├── parser.go        # ParseBoard desde layout texto
│   └── solver.go        # Backtracking determinista
├── dictionary/
│   └── dictionary.go    # LoadDictionary, normalización (Ñ, sin tildes)
├── input/
│   ├── board1.txt       # Única plantilla
│   └── words.txt        # Diccionario de ejemplo
└── output/              # Gitignored; cruzada.html
```

### Comportamientos conocidos

- El solver prueba candidatos en **orden fijo** del diccionario → misma solución en cada run
- Salida solo **HTML** embebido en `main.go`
- Paths **hardcodeados** (`input/board1.txt`, etc.)
- **Sin tests**, **sin CI**
- Longitud mínima de slot: **4 runes** (regla de negocio en parser)

### Dependencias

- Go module: `github.com/matiasnoriega/mechigram`
- Sin dependencias externas hoy (stdlib only)

## Estado objetivo (Iteración 2 — planificado)

Ver [`ROADMAP.md`](ROADMAP.md). Resumen:

```
cmd/mechigram/           # CLI con flags
internal/app/            # Orquestación Generate()
core/                    # Dominio puro (sin cambios de responsabilidad)
dictionary/
render/
├── pdf/
└── html/
templates/ o input/boards/
```

## Restricciones no negociables

1. **`core/` no importa** `os`, HTTP, PDF, ni paquetes de render
2. **Soporte de Ñ** y normalización sin tildes en diccionario y tablero
3. **PlaceWord / RemoveWord** con recibo (`[]bool`) — no borrar letras de otras palabras
4. **Impresión** es el criterio de éxito del render, no la preview en pantalla
5. **Scope mínimo** — no implementar features de Iteración 3 salvo petición explícita

## Usuarios del sistema

| Actor | Necesidad |
|-------|-----------|
| Abuela (jugadora) | PDF claro, imprimible, palabras listadas |
| Maintainer (tú) | Generar variedad, batch futuro, mantener OSS |
| Contribuidor externo | Docs, tests, convenciones claras |
| Agente IA | Harness en `ia/` para no romper arquitectura |

## Referencias externas

- Formato visual inspirado en revista **33 Cruzadas** (Edigrama) — validar con impresión real
- Fill-in crossword: todas las palabras dadas; el jugador las coloca en el grid
