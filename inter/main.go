func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := make(map[rune]bool)

	for _, r1 := range s1 {
		if seen[r1] {
			continue
		}

		for _, r2 := range s2 {
			if r1 == r2 {
				z01.PrintRune(r1)
				seen[r1] = true
				break
			}
		}
	}

	z01.PrintRune('\n')
}