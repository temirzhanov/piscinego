package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	}

	if len(slice1) == 0 {
		return slice2
	}
	if len(slice2) == 0 {
		return slice1
	}

	for _, v := range slice2 {
		slice1 = append(slice1, v)
	}
	return slice1
}
