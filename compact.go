package piscine

func Compact(ptr *[]string) int {
	count := 0

	for _, value := range *ptr {
		if value != "" {
			count++
		}
	}
	res := make([]string, count)
	ind := 0
	for _, value := range *ptr {
		if value != "" {
			res[ind] = value
			ind++
		}
	}
	*ptr = res
	return count
}
