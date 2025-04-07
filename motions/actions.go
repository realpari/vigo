package motions

import (
	"strconv"
	"unicode"
)

func BasicActions(keys string) []string {
	var actions []string
	var countRune string

	for _, key := range keys {
		if unicode.IsDigit(key) {
			countRune = countRune + string(key)

		} else if unicode.IsLetter(key) {

			if countRune == "" {
				countRune = "1"
			}

			count, _ := strconv.Atoi(countRune)

			var action string
			switch key {
			case 'd':
				action = "delete"
			case 'y':
				action = "yank"
			case 'p':
				action = "paste"
			case 'f':
				action = "find"
			default:
				action = "unknown key"
			}

			for i := 0; i < count; i++ {
				actions = append(actions, action)
			}
			countRune = ""
		}
	}

	return actions
}
