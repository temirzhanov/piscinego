package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	lenA := lenArr(a)

	for i := 0; i < lenA-1; i++ {
		for j := i + 1; j < lenA; j++ {
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
