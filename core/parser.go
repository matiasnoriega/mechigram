package core

import "strings"

// ParseBoard toma un array de strings (el dibujo del tablero) y devuelve un Board listo.
func ParseBoard(layout []string) *Board {
	if len(layout) == 0 {
		return nil
	}

	height := len(layout)
	width := len(layout[0])

	// 1. Inicializamos la matriz bidimensional (Grid)
	grid := make([][]Cell, height)
	for y, row := range layout {
		grid[y] = make([]Cell, width)
		runes := []rune(strings.ToUpper(row))

		for x := 0; x < width; x++ {
			if x < len(runes) && runes[x] == '#' {
				grid[y][x] = Cell{Char: '#'} // Bloque negro
			} else {
				grid[y][x] = Cell{Char: ' '} // Espacio en blanco libre
			}
		}
	}

	var slots []Slot
	slotID := 1

	// 2. Detectar Slots Horizontales
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Es el inicio de una palabra si está vacío Y (está en el borde izquierdo O tiene un '#' a su izquierda)
			if grid[y][x].Char == ' ' && (x == 0 || grid[y][x-1].Char == '#') {
				length := 0
				// Medimos cuántos espacios en blanco hay hacia la derecha
				for currX := x; currX < width && grid[y][currX].Char == ' '; currX++ {
					length++
				}
				// En las cruzadas reales no hay palabras de 1 sola letra
				if length >= 4 {
					slots = append(slots, Slot{
						ID:        slotID,
						StartX:    x,
						StartY:    y,
						Length:    length,
						Direction: Horizontal,
					})
					slotID++
				}
			}
		}
	}

	// 3. Detectar Slots Verticales
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Es el inicio de una palabra si está vacío Y (está en el borde superior O tiene un '#' arriba)
			if grid[y][x].Char == ' ' && (y == 0 || grid[y-1][x].Char == '#') {
				length := 0
				// Medimos cuántos espacios en blanco hay hacia abajo
				for currY := y; currY < height && grid[currY][x].Char == ' '; currY++ {
					length++
				}
				if length >= 4 {
					slots = append(slots, Slot{
						ID:        slotID,
						StartX:    x,
						StartY:    y,
						Length:    length,
						Direction: Vertical,
					})
					slotID++
				}
			}
		}
	}

	return &Board{
		Width:  width,
		Height: height,
		Grid:   grid,
		Slots:  slots,
	}
}
