package piscine

func Unmatch(a []int) int {
	for i, iNumb := range a {
		count := 1
		for j, jNumb := range a {
			if iNumb == jNumb && i != j {
				count++
			}
		}

		if count%2 == 1 {
			return iNumb
		}
	}
	return -1
}
