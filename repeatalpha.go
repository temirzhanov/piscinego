package piscine

func RepeatAlpha(s string) string {
	res := ""

	for _, ch := range s {
		diff := 1
		if ch >= 'a' && ch <= 'z' {
			diff = int(ch-'a') + 1
		} else if ch >= 'A' && ch <= 'Z' {
			diff = int(ch-'A') + 1
		}

		for i := 0; i < diff; i++ {
			res += string(ch)
		}
	}

	return res
}
