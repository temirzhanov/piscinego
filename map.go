package piscine

func Map(f func(int) bool, a []int) []bool {
	lenA := lenArrInt(a)

	res := make([]bool, lenA)

	for ind := range a {
		res[ind] = f(a[ind])
	}

	return res
}
