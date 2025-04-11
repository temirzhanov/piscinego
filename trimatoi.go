package piscine

func TrimAtoi(s string) int {
	numb := 0
	sign := 1
	check := true
	for _, letter := range s {
		if letter == '-' && check {
			sign *= -1
		}
		if letter >= '0' && letter <= '9' {
			numb = numb*10 + int(letter-48)
			check = false
		}
	}

	return numb * sign
}
