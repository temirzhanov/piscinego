package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	res := ""
	for _, letter := range args[0] {
		if string(letter) == args[1] {
			res += args[2]
		} else {
			res += string(letter)
		}
	}

	printString(res)
}

func printString(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
