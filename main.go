package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/matiasnoriega/mechigram/core"
	"github.com/matiasnoriega/mechigram/dictionary"
)

func main() {
	// 1. Cargar el tablero
	boardData, err := os.ReadFile("input/board1.txt")
	if err != nil {
		log.Fatalf("Error leyendo el tablero: %v", err)
	}

	lines := strings.Split(strings.ReplaceAll(string(boardData), "\r\n", "\n"), "\n")
	var boardLines []string
	for _, l := range lines {
		if l != "" {
			boardLines = append(boardLines, l)
		}
	}

	board := core.ParseBoard(boardLines)
	if board == nil {
		log.Fatal("Error al parsear el tablero")
	}

	// 2. Cargar el diccionario
	dict, err := dictionary.LoadDictionary("input/words.txt")
	if err != nil {
		log.Fatalf("Error leyendo el diccionario: %v", err)
	}

	// 3. Resolver
	solver := core.NewSolver(board, dict)
	if solver.Solve() {
		fmt.Println("¡Solución encontrada!")
		renderHTML(board, solver.GetPlacedWords(), "output/cruzada.html")
		fmt.Println("HTML generado en output/cruzada.html")
	} else {
		fmt.Println("No se pudo encontrar una solución con las palabras dadas.")
	}
}

func renderHTML(b *core.Board, words []string, outputPath string) {
	// Crear el directorio output si no existe
	if err := os.MkdirAll("output", 0755); err != nil {
		log.Fatalf("Error creando directorio output: %v", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creando archivo HTML: %v", err)
	}
	defer file.Close()

	html := `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Cruzada Mechigram</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #ffffff;
            color: #000000;
            display: flex;
            flex-direction: column;
            align-items: center;
            margin: 0;
            padding: 40px;
        }
        h1 {
            font-size: 3em;
            margin-bottom: 20px;
        }
        .grid {
            display: grid;
            grid-template-columns: repeat(` + fmt.Sprint(b.Width) + `, 60px);
            grid-template-rows: repeat(` + fmt.Sprint(b.Height) + `, 60px);
            gap: 2px;
            background-color: #000;
            border: 4px solid #000;
            margin-bottom: 40px;
        }
        .cell {
            background-color: #fff;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 2.5em;
            font-weight: bold;
            text-transform: uppercase;
        }
        .cell.black {
            background-color: #000;
        }
        .words-container {
            width: 100%;
            max-width: 800px;
        }
        .words-title {
            font-size: 2em;
            border-bottom: 3px solid #000;
            padding-bottom: 10px;
            margin-bottom: 20px;
        }
        .words-list {
            display: flex;
            flex-wrap: wrap;
            gap: 15px;
            list-style: none;
            padding: 0;
        }
        .words-list li {
            font-size: 1.8em;
            padding: 5px 15px;
            border: 2px solid #000;
            border-radius: 8px;
        }
    </style>
</head>
<body>
    <h1>Cruzada</h1>
    <div class="grid">`

	// Renderizar el tablero
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			c := b.Grid[y][x].Char
			if c == '#' {
				html += `<div class="cell black"></div>`
			} else {
				html += fmt.Sprintf(`<div class="cell">%c</div>`, c)
			}
		}
	}

	html += `
    </div>
    <div class="words-container">
        <div class="words-title">Palabras a buscar:</div>
        <ul class="words-list">`

	// Renderizar las palabras
	for _, word := range words {
		html += fmt.Sprintf(`<li>%s</li>`, word)
	}

	html += `
        </ul>
    </div>
</body>
</html>`

	_, err = file.WriteString(html)
	if err != nil {
		log.Fatalf("Error escribiendo HTML: %v", err)
	}
}
