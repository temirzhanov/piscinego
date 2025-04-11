package piscine

func ForEach(f func(int), a []int) {
	for ind := range a {
		f(a[ind])
	}
}
