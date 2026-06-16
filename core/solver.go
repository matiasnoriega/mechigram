package core

// Solver maneja el estado del algoritmo de resolución.
type Solver struct {
	board       *Board
	dict        map[int][]string // Diccionario de palabras agrupadas por longitud
	usedWords   map[string]bool  // Palabras ya usadas en el tablero
	wordsPlaced []string         // Palabras ubicadas, en orden de slot
}

// NewSolver inicializa un nuevo Solver para el tablero dado.
func NewSolver(b *Board, dictionary map[int][]string) *Solver {
	return &Solver{
		board:       b,
		dict:        dictionary,
		usedWords:   make(map[string]bool),
		wordsPlaced: make([]string, len(b.Slots)),
	}
}

// Solve ejecuta el algoritmo de backtracking.
// Devuelve true si encuentra una solución válida, false de lo contrario.
func (s *Solver) Solve() bool {
	return s.solveRecursive(0)
}

// solveRecursive es la función interna recursiva que recorre los slots.
func (s *Solver) solveRecursive(slotIndex int) bool {
	// Si ya llenamos todos los slots, terminamos con éxito
	if slotIndex == len(s.board.Slots) {
		return true
	}

	slot := s.board.Slots[slotIndex]
	candidates := s.dict[slot.Length]

	// Probar cada palabra que tenga la longitud correcta
	for _, word := range candidates {
		// Si la palabra ya fue usada, la salteamos
		if s.usedWords[word] {
			continue
		}

		// Verificar si la palabra encaja en el slot sin conflictos
		if s.board.CanPlaceWord(word, slot) {
			// Intentar colocar la palabra
			placed := s.board.PlaceWord(word, slot)
			s.usedWords[word] = true
			s.wordsPlaced[slotIndex] = word

			// Llamada recursiva al siguiente slot
			if s.solveRecursive(slotIndex + 1) {
				return true // Solución encontrada en esta rama
			}

			// Backtracking: deshacer cambios si no se encontró solución
			s.board.RemoveWord(slot, placed)
			s.usedWords[word] = false
			s.wordsPlaced[slotIndex] = ""
		}
	}

	// Si ninguna palabra sirvió para este slot, retornamos false para retroceder
	return false
}

// GetPlacedWords retorna la lista de palabras usadas, útil para mostrar debajo del tablero.
func (s *Solver) GetPlacedWords() []string {
	var result []string
	for _, w := range s.wordsPlaced {
		if w != "" {
			result = append(result, w)
		}
	}
	return result
}
