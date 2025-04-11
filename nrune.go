package piscine

func NRune(s string, n int) rune {
	for ind, letter := range s {
		if n == ind+1 {
			return letter
		}
	}
	return '\x00'
}
