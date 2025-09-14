package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	for i, arg := range os.Args[1:] {
		runes := []rune(arg)
		length := len(runes)

		for j := 0; j < length; j++ {
			r := runes[j]

			if r >= 'A' && r <= 'Z' {
				r += 32
			}

			isLastLetter := (j == length-1) ||
				(runes[j+1] == ' ') || (runes[j+1] == '\t')

			if isLastLetter && r >= 'a' && r <= 'z' {
				r -= 32
			}

			z01.PrintRune(r)
		}

		z01.PrintRune('\n')

		_ = i
	}
}
