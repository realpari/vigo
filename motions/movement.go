package motions

func BasicMovement(keys string) string {
	var movement string

	switch keys {
	case "k":
		movement = "up"
	case "j":
		movement = "down"
	case "h":
		movement = "left"
	case "l":
		movement = "right"
	default:
		movement = "uknown key"
	}

	return movement
}
