package piscine

func CheckNumber(arg string) bool {
	for _, ch := range arg {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}
