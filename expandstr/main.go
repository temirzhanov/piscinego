package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	space := false
	printed := false

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			space = true
		} else {
			if space && printed {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(s[i]))
			space = false
			printed = true
		}
	}
	z01.PrintRune('\n')
}
