package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}
	var res []int
	for i := max; i > min; i-- {
		res = append(res, i)
	}

	return res
}
