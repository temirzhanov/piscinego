package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	for i, b := range arr {
		z01.PrintRune(rune(hex[b>>4]))
		z01.PrintRune(rune(hex[b&0x0F]))
		linePos := i % 4
		isLastInLine := linePos == 3 || i == len(arr)-1
		if !isLastInLine {
			z01.PrintRune(' ')
		}

		if isLastInLine {
			z01.PrintRune('\n')
		}
	}
	for _, b := range arr {
		r := rune(b)
		if r >= 33 && r <= 126 {
			z01.PrintRune(r)
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
