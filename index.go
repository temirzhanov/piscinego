package piscine

func Index(s string, toFind string) int {
	lenS := StrLen(s)
	lenToFind := StrLen(toFind)

	if lenS < lenToFind {
		return -1
	}

	for ind := range s {
		check := true
		for indX := range toFind {
			// fmt.Println("checker")
			// fmt.Println(ind)
			if ind+indX == lenS {
				return -1
			}
			// fmt.Println(s[ind+indX])
			// fmt.Println(toFind[indX])
			if s[ind+indX] != toFind[indX] {
				check = false
				break
			}
		}
		if check {
			return ind
		}
	}

	return -1
}
