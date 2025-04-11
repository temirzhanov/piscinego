package piscine

func JumpOver(str string) string {
	res := ""
	for ind, letter := range str {
		if (ind+1)%3 == 0 {
			res += string(letter)
		}
	}

	return res + "\n"
}
