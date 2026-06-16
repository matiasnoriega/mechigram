package dictionary

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

// LoadDictionary lee un archivo de texto, procesa cada palabra y
// devuelve un mapa agrupado por longitud de las runas.
func LoadDictionary(filename string) (map[int][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dict := make(map[int][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawWord := strings.TrimSpace(scanner.Text())
		if rawWord == "" {
			continue
		}

		cleanWord := processWord(rawWord)
		// Ignoramos palabras vacías después de procesar
		if cleanWord == "" {
			continue
		}

		length := len([]rune(cleanWord))
		dict[length] = append(dict[length], cleanWord)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dict, nil
}

// processWord normaliza la palabra: quita tildes, convierte a mayúsculas,
// manteniendo la 'Ñ' intacta.
func processWord(word string) string {
	// 1. Convertir a mayúsculas
	word = strings.ToUpper(word)

	// 2. Remover tildes usando norm y runes.Remove (dejando la Ñ tranquila si se puede,
	// pero golang/x/text es un poco complejo para la Ñ. Una alternativa sencilla es un reemplazo directo:
	replacer := strings.NewReplacer(
		"Á", "A",
		"É", "E",
		"Í", "I",
		"Ó", "O",
		"Ú", "U",
		"Ü", "U",
	)
	word = replacer.Replace(word)

	// 3. Filtrar cualquier carácter que no sea letra (como signos de puntuación)
	var builder strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
