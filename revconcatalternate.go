package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	res := make([]int, 0, len(slice1)+len(slice2))

	i := len(slice1) - 1
	j := len(slice2) - 1

	for i > j {
		res = append(res, slice1[i])
		i--
	}
	for j > i {
		res = append(res, slice2[j])
		j--
	}

	first := true

	for i >= 0 && j >= 0 {
		if first {
			res = append(res, slice1[i])
			i--
		} else {
			res = append(res, slice2[j])
			j--
		}
		first = !first
	}

	for i >= 0 {
		res = append(res, slice1[i])
		i--
	}
	for j >= 0 {
		res = append(res, slice2[j])
		j--
	}

	return res
}
