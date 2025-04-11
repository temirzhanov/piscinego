package piscine

func ToLower(s string) string {
	str := ""
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			str += string(letter + 32)
			continue
		}
		str += string(letter)
	}

	return str
}
