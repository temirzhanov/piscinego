package piscine

func SortIntegerTable(table []int) {
	length := lenArrInt(table)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func lenArrInt(table []int) int {
	count := 0
	for range table {
		count++
	}
	return count
}
