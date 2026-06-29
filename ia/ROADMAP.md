# Roadmap — Mechigram

Plan de evolución por iteraciones. **Iteración 2** es el foco activo.

## Principios de priorización

| Criterio | Peso | Implicación |
|----------|------|-------------|
| Valor usuaria final (impresión) | Alto | PDF nativo primero |
| Calidad portfolio OSS | Alto | Tests + CI + capas en Iter. 2 |
| Complejidad / riesgo | Medio | Concurrencia → Iter. 3 |
| Deuda técnica | Medio | Refactor CLI y RNG en Iter. 2 |

**Decisión:** Iteración 2 entrega **una cruzada PDF imprimible**, **variabilidad real** y **cimientos OSS**. Generación masiva en paralelo queda para Iteración 3.

---

## Iteración 1 — Completada ✓

- [x] `core`: Board, Slot, Parser, PlaceWord/RemoveWord con recibo
- [x] Solver backtracking
- [x] Diccionario desde archivo con normalización ES
- [x] HTML rudimentario

---

## Iteración 2 — En planificación (objetivo activo)

> Generar una cruzada en **PDF listo para imprimir** (A4, tipografía grande), con **palabras distintas** en cada ejecución, vía **CLI**, con **tests en dominio** y **CI**.

### Épica 1 — Arquitectura y CLI

| ID | Tarea | Entregable |
|----|-------|------------|
| 1.1 | Entrypoint `cmd/mechigram/main.go` | `main` solo orquesta |
| 1.2 | Paquete `internal/app` | `Generate(opts) → Result` |
| 1.3 | Interfaz `Renderer` | `Render(Puzzle) error` |
| 1.4 | Flags CLI | `--board`, `--dict`, `--output`, `--seed`, `--format` |
| 1.5 | Tipo `Puzzle` | Tablero + palabras ordenadas alfabéticamente |

### Épica 2 — PDF nativo

| ID | Tarea | Notas |
|----|-------|-------|
| 2.1 | Librería PDF | Preferencia: `gofpdf`; TTF embebida para Ñ |
| 2.2 | `render/pdf` | Independiente del solver |
| 2.3 | Layout tipo 33 Cruzadas | A4, celdas ~12–14 mm, lista en columnas |
| 2.4 | `--solution` | Puzzle vacío vs con letras |
| 2.5 | `render/html` | Fallback dev; extraer de `main.go` |
| 2.6 | Validación impresión | Prueba en impresora real |

**DoD PDF:** Legible a ~50 cm, sin recortes, palabras en misma página (máx. 2 si tablero grande).

### Épica 3 — Variabilidad

| ID | Tarea |
|----|-------|
| 3.1 | RNG inyectado en `Solver` |
| 3.2 | Shuffle candidatos (Fisher-Yates) |
| 3.3 | `--seed` (0 = aleatorio) |
| 3.4 | Reintentos ante fallo (N veces) |
| 3.5 | Cargador multi-plantilla |
| 3.6 | Selección aleatoria de plantilla |
| 3.7 | ≥ 3 plantillas en `input/boards/` |

### Épica 4 — Tests

| ID | Paquete | Meta |
|----|---------|------|
| 4.1 | `core/board` | Place/Remove, choques |
| 4.2 | `core/parser` | Slots H/V, min length 4 |
| 4.3 | `core/solver` | Fixture + seed fijo |
| 4.4 | `dictionary` | Normalización Ñ/tildes |
| 4.5 | `render/pdf` | Smoke `%PDF` |
| 4.6 | Golden (opc.) | Hash PDF con seed fijo |

**Meta cobertura:** ≥ 80 % en `core` y `dictionary`.

### Épica 5 — CI/CD y OSS

| ID | Tarea |
|----|-------|
| 5.1 | `.github/workflows/ci.yml` |
| 5.2 | Matriz Go 1.22+ |
| 5.3 | golangci-lint (opcional) |
| 5.4 | README ampliado |
| 5.5 | CONTRIBUTING.md |
| 5.6 | Issue templates |

### Orden sugerido

1. Semana 1: Épica 1 + tests board/parser
2. Semana 2: Épica 2 (PDF)
3. Semana 3: Épica 3 (RNG + plantillas)
4. Semana 4: Épica 5 + pulido

### Definición de Done — Iteración 2

- [ ] CLI genera PDF modo puzzle + lista de palabras
- [ ] Dos runs consecutivos → palabras diferentes (misma plantilla)
- [ ] ≥ 3 plantillas disponibles
- [ ] `go test ./...` local + CI verde
- [ ] README con arquitectura y flags
- [ ] `core` sin imports de I/O ni render

---

## Iteración 3 — Preview (no ejecutar aún)

| Épica | Contenido |
|-------|-----------|
| Lote | Worker pool, N puzzles → PDF multipágina |
| Solver avanzado | MRV, timeout, métricas |
| Diccionario | Wordlist ES externo, filtros frecuencia |
| API | REST / Lambda + S3 |
| Revista | 33 cruzadas, portada, numeración |

---

## Riesgos

| Riesgo | Mitigación |
|--------|------------|
| Fuente PDF sin Ñ | TTF embebida (DejaVu / Noto) |
| Sin solución tras shuffle | Reintentos + mensaje claro |
| Layout no cabe A4 | Escalar celdas por dimensiones |
| Scope creep | Concurrencia explícitamente fuera de Iter. 2 |
