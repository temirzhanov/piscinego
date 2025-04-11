package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, word := range args {
		runes := []rune(word)
		for ind := range word {
			z01.PrintRune(runes[ind])
		}
		z01.PrintRune('\n')
	}
}
