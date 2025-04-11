package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if checkLen(args) == 1 {
		return
	}
	args = args[1:]

	div := 96
	if args[0] == "--upper" {
		div = 64
		args = args[1:]
	}

	for _, number := range args {
		numb := Atoi(number) + div
		if (numb >= 65 && numb <= 90) || (numb >= 97 && numb <= 122) {
			z01.PrintRune(rune(numb))
		} else {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')
}

func checkLen(strArr []string) int {
	count := 0

	for range strArr {
		count++
	}

	return count
}

func Atoi(s string) int {
	number := 0
	check := true
	sign := 1

	for ind, let := range s {
		if let == '-' && check && ind == 0 {
			sign = -1
			check = false
		} else if let == '+' && check && ind == 0 {
			check = false
		} else if let >= '0' && let <= '9' {
			number = number*10 + int(let) - 48
		} else {
			return 0
		}
	}

	return number * sign
}
