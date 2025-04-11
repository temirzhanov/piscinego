package piscine

func StringToIntSlice(str string) []int {
	bt := []rune(str)
	var res []int
	for _, numb := range bt {
		res = append(res, int(numb))
	}

	return res
}
