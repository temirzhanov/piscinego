package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	s1 := args[0]
	s2 := args[1]

	if len(s1) == 0 {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	i := 0 // позиция в r2
	for _, ch := range r1 {
		found := false
		for i < len(r2) {
			if r2[i] == ch {
				i++
				found = true
				break
			}
			i++
		}
		if !found {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	z01.PrintRune('1')
	z01.PrintRune('\n')
}
