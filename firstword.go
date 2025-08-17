package piscine

func FirstWord(s string) string {
	word := ""
	check := false

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			check = true
			word += string(char)
		} else if check {
			break
		}
	}

	return word + "\n"
}
