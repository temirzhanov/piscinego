package piscine

func PrintIf(str string) string {
	if len(str) > 2 || len(str) == 0 {
		return "G\n"
	}
	return "Invalid Input\n"
}
