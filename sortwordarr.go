package piscine

func SortWordArr(a []string) {
	lenA := lenArr(a)

	for i := 0; i < lenA-1; i++ {
		for j := i + 1; j < lenA; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	// fmt.Println(a)
}
