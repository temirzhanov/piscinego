package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	argLen := checkLen(args)
	if argLen == 0 {
		return
	}

	for i := argLen - 1; i >= 0; i-- {
		runes := []rune(args[i])
		for ind := range args[i] {
			z01.PrintRune(runes[ind])
		}
		z01.PrintRune('\n')
	}
}

func checkLen(str []string) int {
	counter := 0
	for range str {
		counter++
	}

	return counter
}
