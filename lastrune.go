package piscine

func LastRune(s string) rune {
	counter := -1
	for range s {
		counter++
	}
	sent := []rune(s)
	return sent[counter]
}
