package piscine

func CountAlpha(s string) int {

	count := 0
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			count++
		}
	}

	return count

}
