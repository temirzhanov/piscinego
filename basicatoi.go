package piscine

func BasicAtoi(s string) int {
	number := 0

	for _, let := range s {
		if let == '-' {
			number *= -1
		} else if let >= '0' && let <= '9' {
			number = number*10 + int(let) - 48
		}
	}

	return number
}
