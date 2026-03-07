package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	str := os.Args[1]
	inWord := false
	printedWord := false

	for _, r := range str {
		if r != ' ' && r != '\t' {
			if !inWord {
				if printedWord {
					z01.PrintRune(' ')
				}
				inWord = true
				printedWord = true
			}
			z01.PrintRune(r)
		} else {
			inWord = false
		}
	}

	z01.PrintRune('\n')

}
