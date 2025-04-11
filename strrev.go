package piscine

func StrRev(s string) string {
	rev := ""

	for _, let := range s {
		rev = string(let) + rev
	}

	return rev
}
