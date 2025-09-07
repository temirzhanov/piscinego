package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func isSpace(b byte) bool {
	return b == ' ' || b == '\t'
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	s := os.Args[1]
	tokens := []string{}
	start := -1

	for i := 0; i < len(s); i++ {
		if !isSpace(s[i]) {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				tokens = append(tokens, s[start:i])
				start = -1
			}
		}
	}
	if start != -1 {
		tokens = append(tokens, s[start:])
	}

	if len(tokens) == 0 {
		z01.PrintRune('\n')
		return
	}

	if len(tokens) > 1 {
		first := tokens[0]
		tokens = append(tokens[1:], first)
	}

	for i, t := range tokens {
		if i > 0 {
			z01.PrintRune(' ')
		}
		printString(t)
	}
	z01.PrintRune('\n')
}
