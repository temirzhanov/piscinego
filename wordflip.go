package piscine

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}

	// Trim spaces manually
	runes := []rune(str)
	start, end := 0, len(runes)-1
	for start <= end && runes[start] == ' ' {
		start++
	}
	for end >= start && runes[end] == ' ' {
		end--
	}

	// Case: whitespace-only input
	if start > end {
		return "\n"
	}

	// Split words manually (ignoring multiple spaces)
	words := []string{}
	word := ""
	for i := start; i <= end; i++ {
		if runes[i] == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(runes[i])
		}
	}
	if word != "" {
		words = append(words, word)
	}

	// Reverse words
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		if i > 0 {
			res += " "
		}
	}

	return res + "\n"
}
