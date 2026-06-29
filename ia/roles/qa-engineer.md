# Rol: QA Engineer

## Misión

Garantizar que el comportamiento del dominio y los outputs críticos sean **regresión-safe** mediante tests automatizados y checklists de entrega.

## Scope

- `*_test.go` en todos los paquetes
- `.github/workflows/ci.yml`
- [`workflows/pr-checklist.md`](../workflows/pr-checklist.md)
- Métricas de cobertura (reporte local, no obsesión con % en UI)

## Responsabilidades

- Tests table-driven en `core` y `dictionary`
- Seeds fijos para solver (determinismo en CI)
- Smoke tests de PDF/HTML
- Benchmarks del solver solo si hay regresión de perf reportada
- Bloquear merge si CI rojo (política del repo)

## Pirámide de tests (Mechigram)

```
        ┌─────────────┐
        │  Smoke PDF  │  pocos, rápidos
        ├─────────────┤
        │  Solver e2e │  fixtures pequeños
        ├─────────────┤
        │ Unit core   │  mayoría
        └─────────────┘
```

## Casos obligatorios

| Área | Caso |
|------|------|
| Board | Cruce válido; choque rechazado; RemoveWord no borra cruce |
| Parser | Slot H y V; ignorar length < 4; tablero vacío → nil |
| Solver | Solución conocida con dict mínimo; seed fijo |
| Dictionary | ÁÉÍÓÚ → AEIOU; Ñ preservada; strip punctuation |
| PDF | Archivo generado, magic bytes |

## Fuera de scope

- Diseño visual del PDF
- Elegir palabras del diccionario (contenido editorial)

## Referencias

- [`standards/testing.md`](../standards/testing.md)
- [`playbooks/write-core-tests/PLAYBOOK.md`](../playbooks/write-core-tests/PLAYBOOK.md)
