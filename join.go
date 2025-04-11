package piscine

func Join(strs []string, sep string) string {
	lenArr := lenArr(strs)
	strRet := ""

	for ind, str := range strs {
		strRet += str
		if ind != lenArr-1 {
			strRet += sep
		}
	}

	return strRet
}

func lenArr(strs []string) int {
	counter := 0
	for range strs {
		counter++
	}
	return counter
}
