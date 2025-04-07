package motions

func basicMotion(key string) string {
	switch key {
	case "j":
		return "down"
	case "k":
		return "up"
	case "h":
		return "left"
	case "l":
		return "right"
	default:
		return "error"
	}
}
