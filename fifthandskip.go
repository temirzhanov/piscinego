package piscine

func FifthAndSkip(str string) string {

	res := ""
	counter := 0
	if str == "" {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Input\n"
	}

	for _, char := range str {
		if char != ' ' {
			counter++
			if counter > 5 {
				counter = 0
				res += " "
			} else {
				res += string(char)
			}
		}
	}

	return res
}
