package piscine

func LastWord(s string) string {
	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}
	if i < 0 {
		return "\n"
	}

	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}

	return s[i+1:end+1] + "\n"
}
