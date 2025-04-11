package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for ind := range tab {
		if f(tab[ind]) {
			count++
		}
	}

	return count
}
