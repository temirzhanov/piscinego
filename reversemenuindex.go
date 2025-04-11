package piscine

func ReverseMenuIndex(menu []string) []string {
	lenA := len(menu)

	res := make([]string, lenA)

	for i := 0; i < lenA; i++ {
		res[lenA-1-i] = menu[i]
	}

	return res
}
