package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	length := lenArrInt(a)

	asciending := true
	descending := true

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			descending = false
		}
	}

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			asciending = false
		}
	}

	return asciending || descending
}
