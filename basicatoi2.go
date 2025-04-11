package piscine

func BasicAtoi2(s string) int {
	number := 0

	for _, let := range s {
		if let == '-' {
			number *= -1
		} else if let >= '0' && let <= '9' {
			number = number*10 + int(let) - 48
		} else {
			return 0
		}
	}

	return number
}
