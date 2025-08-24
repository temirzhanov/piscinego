package piscine

func HashCode(dec string) string {
	res := ""

	for _, c := range dec {
		num := (int(c) + len(dec)) % 127
		if num < 33 {
			num += 33
		}
		res += string(rune(num))
	}

	return res
}
