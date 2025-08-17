package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	} else if len(slice1) == 0 {
		return slice2
	} else if len(slice2) == 0 {
		return slice1
	}

	first := slice1
	second := slice2

	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	}

	var res []int

	for i := 0; i < len(first); i++ {
		res = append(res, first[i])

		if i < len(second) {
			res = append(res, second[i])
		}
	}

	return res
}
