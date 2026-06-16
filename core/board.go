package core

// Direction define si el hueco para la palabra es horizontal o vertical.
type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

// Cell representa una celda individual en el tablero.
type Cell struct {
	Char rune // La letra actual (' ' si está vacía, '#' si es un bloque negro).
	// Nota: Usamos `rune` para manejar bien caracteres en español como la 'Ñ'.
}

// Slot representa un "hueco" donde debe encajar una palabra exacta de la lista.
type Slot struct {
	ID        int
	StartX    int       // Columna de inicio.
	StartY    int       // Fila de inicio.
	Length    int       // Cantidad de letras que entran en este hueco.
	Direction Direction // Horizontal o Vertical.
}

// Board es el estado global de la cruzada en un momento dado.
type Board struct {
	Width  int
	Height int
	Grid   [][]Cell // Matriz bidimensional (Y, X) para acceso O(1) a cualquier celda.
	Slots  []Slot   // Lista de todos los espacios que el algoritmo debe llenar.
}

// CanPlaceWord verifica si una palabra encaja en el slot sin chocar con letras existentes.
func (b *Board) CanPlaceWord(word string, slot Slot) bool {
	runes := []rune(word) // Convertimos a runas para soportar tildes o la 'Ñ'

	// Por seguridad, validamos que la longitud sea exacta
	if len(runes) != slot.Length {
		return false
	}

	for i := 0; i < slot.Length; i++ {
		x, y := slot.StartX, slot.StartY
		if slot.Direction == Horizontal {
			x += i
		} else {
			y += i
		}

		currentCell := b.Grid[y][x].Char
		// Si la celda no está vacía y la letra NO coincide, hay un choque
		if currentCell != ' ' && currentCell != runes[i] {
			return false
		}
	}
	return true
}

// PlaceWord escribe la palabra en el tablero.
// Devuelve un array de booleanos indicando qué posiciones estaban vacías y fueron llenadas por esta palabra.
func (b *Board) PlaceWord(word string, slot Slot) []bool {
	runes := []rune(word)
	placed := make([]bool, slot.Length)

	for i := 0; i < slot.Length; i++ {
		x, y := slot.StartX, slot.StartY
		if slot.Direction == Horizontal {
			x += i
		} else {
			y += i
		}

		// Si la celda estaba vacía, la escribimos y guardamos en nuestro "recibo" que nosotros la pusimos
		if b.Grid[y][x].Char == ' ' {
			b.Grid[y][x].Char = runes[i]
			placed[i] = true
		}
	}
	return placed
}

// RemoveWord deshace el cambio usando el "recibo" generado por PlaceWord.
// Así nos aseguramos de no borrar letras que pertenecen a otras palabras cruzadas.
func (b *Board) RemoveWord(slot Slot, placed []bool) {
	for i := 0; i < slot.Length; i++ {
		// Solo vaciamos la celda si fuimos nosotros los que pusimos esa letra
		if placed[i] {
			x, y := slot.StartX, slot.StartY
			if slot.Direction == Horizontal {
				x += i
			} else {
				y += i
			}
			b.Grid[y][x].Char = ' '
		}
	}
}
