package piscine

func Capitalize(s string) string {
	str := ""
	check := true
	for _, letter := range s {
		if (letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9') {
			if check {
				if letter >= 'a' && letter <= 'z' {
					str += string(letter - 32)
					check = false
					continue
				}
				str += string(letter)
				check = false
			} else {
				if letter >= 'A' && letter <= 'Z' {
					str += string(letter + 32)
					continue
				}
				str += string(letter)
			}
		} else {
			str += string(letter)
			check = true
		}
	}

	return str
}
