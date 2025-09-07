package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	inWord := false
	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == ' ' || ch == '\t' || ch == '\n' {
			inWord = false
			continue
		}

		if !inWord {
			inWord = true
			if ch >= 'a' && ch <= 'z' {
				return false
			}
		}
	}
	return true
}
