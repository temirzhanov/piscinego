package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	lenA := checkLen(args)

	for i := 0; i < lenA-1; i++ {
		for j := 0; j < lenA-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
	for _, word := range args {
		runes := []rune(word)
		for ind := range word {
			z01.PrintRune(runes[ind])
		}
		z01.PrintRune('\n')
	}
}

func checkLen(strArr []string) int {
	count := 0

	for range strArr {
		count++
	}

	return count
}
