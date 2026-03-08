package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isSpace(c byte) bool {
	return c == ' ' || c == '\t'
}

func print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	s := os.Args[1]
	i := 0
	n := len(s)

	// пропускаем начальные пробелы
	for i < n && isSpace(s[i]) {
		i++
	}

	// если строка пустая
	if i == n {
		z01.PrintRune('\n')
		return
	}

	// находим первое слово
	start := i
	for i < n && !isSpace(s[i]) {
		i++
	}
	first := s[start:i]

	// печатаем остальные слова
	firstPrinted := false

	for i < n {
		for i < n && isSpace(s[i]) {
			i++
		}
		if i < n {
			if firstPrinted {
				z01.PrintRune(' ')
			}
			wordStart := i
			for i < n && !isSpace(s[i]) {
				i++
			}
			print(s[wordStart:i])
			firstPrinted = true
		}
	}

	// если были другие слова — добавляем пробел
	if firstPrinted {
		z01.PrintRune(' ')
	}

	// печатаем первое слово
	print(first)

	z01.PrintRune('\n')
}
