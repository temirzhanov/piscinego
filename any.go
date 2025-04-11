package piscine

func Any(f func(string) bool, a []string) bool {
	for ind := range a {
		if f(a[ind]) {
			return true
		}
	}

	return false
}
