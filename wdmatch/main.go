package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	target := os.Args[1]
	source := os.Args[2]

	i := 0
	for j := 0; j < len(source) && i < len(target); j++ {
		if source[j] == target[i] {
			i++
		}
	}

	if i == len(target) {
		for _, r := range target {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
