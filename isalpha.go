package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if letter > 'z' || (letter < 'a' && letter > 'Z') || (letter < 'A' && letter > '9') || letter < '0' {
			return false
		}
	}
	return true
}
