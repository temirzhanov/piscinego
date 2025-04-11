package piscine

func ToUpper(s string) string {
	str := ""
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			str += string(letter - 32)
			continue
		}
		str += string(letter)
	}

	return str
}
