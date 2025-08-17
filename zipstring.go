package piscine

func ZipString(s string) string {

	if s == "" {
		return s
	}
	res := ""
	counter := 0
	prevChar := rune(s[0])
	for i, char := range s {
		if prevChar == char {
			counter++

		} else {
			res += string(rune(counter+48)) + string(prevChar)
			prevChar = char
			counter = 1

		}
		if i == len(s)-1 {
			res += string(rune(counter+48)) + string(prevChar)
		}
	}

	return res
}
