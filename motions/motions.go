package motions

import (
	"strconv"
	"unicode"
)

func Vigo(keys string) []string {
	var motions []string
	var countRune string

	for i := 0; i < len(keys); i++ {
		key := rune(keys[i])

		if unicode.IsDigit(key) {
			countRune = countRune + string(key)

		} else if unicode.IsLetter(key) {
			if countRune == "" {
				countRune = "1"
			}

			count, _ := strconv.Atoi(countRune)

			if key == 'k' || key == 'j' || key == 'h' || key == 'l' {
				movements := BasicMovement(string(key))
				for _, direction := range movements {
					for j := 0; j < count; j++ {
						motions = append(motions, direction)
					}
				}
			}

			if key == 'd' || key == 'y' || key == 'p' || key == 'f' {
				actions := BasicActions(string(key))
				for _, action := range actions {
					for j := 0; j < count; j++ {
						motions = append(motions, action)
					}
				}
			}

			countRune = ""
		}
	}

	return motions
}
