package piscine

func Atoi(s string) int {
	number := 0
	check := true
	sign := 1

	for ind, let := range s {
		if let == '-' && check && ind == 0 {
			sign = -1
			check = false
		} else if let == '+' && check && ind == 0 {
			check = false
		} else if let >= '0' && let <= '9' {
			number = number*10 + int(let) - 48
		} else {
			return 0
		}
	}

	return number * sign
}
