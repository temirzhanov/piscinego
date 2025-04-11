package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	runes := []rune(arg)
	for i := range runes {
		if runes[i] == '.' || runes[i] == '/' {
			continue
		}
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
