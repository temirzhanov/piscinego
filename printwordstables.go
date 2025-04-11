package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		runes := []rune(word)

		for ind := range runes {
			z01.PrintRune(runes[ind])
		}
		z01.PrintRune('\n')
	}
}
