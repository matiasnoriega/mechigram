---
name: mechigram-solve-puzzle
description: Implementa o modifica el solver de backtracking, RNG y reintentos en Mechigram. Usar cuando se trabaje en core/solver.go, variabilidad de soluciones, shuffle de candidatos, o flags --seed.
---

# Playbook: Solve Puzzle

## Precondiciones

- Leer `ia/roles/domain-engineer.md`
- Confirmar invariantes de `PlaceWord`/`RemoveWord` en `core/board.go`

## Workflow

### 1. Diagnóstico

- [ ] ¿Cambio es determinismo (tests) o aleatoriedad (producto)?
- [ ] ¿El diccionario tiene suficientes palabras por longitud de slot?

### 2. Implementar RNG (si aplica)

```go
type Solver struct {
    // ...
    rng *rand.Rand
}

func NewSolver(b *Board, dict map[int][]string, rng *rand.Rand) *Solver {
    if rng == nil {
        rng = rand.New(rand.NewSource(time.Now().UnixNano()))
    }
    // ...
}
```

### 3. Shuffle candidatos

En `solveRecursive`, antes del `for _, word := range candidates`:

```go
candidates := append([]string(nil), s.dict[slot.Length]...)
s.shuffleStrings(candidates)
```

Usar Fisher-Yates con `s.rng`.

### 4. Reintentos (capa app, no solver puro)

El solver devuelve bool. Reintentos con nueva semilla viven en `internal/app`:

```go
for attempt := 0; attempt < maxAttempts; attempt++ {
    board := templateBoard.Clone() // o re-parse
    if solver.Solve() { return result, nil }
}
return nil, ErrNoSolution
```

### 5. Tests

- Seed fijo → resultado reproducible
- Dos seeds distintos en fixture con dict amplio → palabras distintas (probabilístico: usar dict mínimo diseñado)

## Verificación

```bash
go test ./core/... -run Solver -v
```

## Referencias

- `ia/playbooks/write-core-tests/PLAYBOOK.md`
- `ia/ROADMAP.md` — Épica 3
