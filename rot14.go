package piscine

func Rot14(s string) string {
	str := ""

	for _, letter := range s {

		if !((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')) {
			str += string(letter)
			continue
		}

		diff := rune(0)
		if letter >= 'A' && letter <= 'Z' {
			diff = 'a' - 'A'
		}

		if letter+diff+14 > 'z' {
			str += string((letter + diff + 14 - 'z') + 'a' - diff - 1)
			continue
		}

		str += string(letter + 14)
	}

	return str
}
