package piscine

func BasicJoin(elems []string) string {
	str := ""

	for _, elem := range elems {
		str += elem
	}

	return str
}
