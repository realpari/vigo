package motions

import (
	"strconv"
	"unicode"
)

func BasicMovement(keys string) []string {
	var movements []string
	var countRune string

	for _, key := range keys {
		if unicode.IsDigit(key) {
			countRune = countRune + string(key)

		} else if unicode.IsLetter(key) {

			if countRune == "" {
				countRune = "1"
			}

			count, _ := strconv.Atoi(countRune)

			var direction string
			switch key {
			case 'k':
				direction = "up"
			case 'j':
				direction = "down"
			case 'h':
				direction = "left"
			case 'l':
				direction = "right"
			default:
				direction = "unknown key"
			}

			for i := 0; i < count; i++ {
				movements = append(movements, direction)
			}
			countRune = ""
		}
	}

	return movements
}
