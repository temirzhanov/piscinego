package piscine

import "github.com/01-edu/z01"

func checkComb(number int) bool {
	for number != 0 {
		if number >= 10 && number%10 <= (number/10)%10 {
			return false
		}
		number /= 10
	}
	return true
}

func printNextDig(numb int) {
	if numb >= 10 {
		printNextDig(numb / 10)
	}
	digit := '0'
	for i := 0; i < numb%10; i++ {
		digit++
	}
	z01.PrintRune(digit)
}

func PrintCombN(n int) {
	maxLength := 1
	for i := 1; i < n; i++ {
		maxLength *= 10
	}
	start := false
	for i := maxLength / 10; i <= maxLength*9; i++ {
		if !checkComb(i) {
			continue
		}
		if start {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		if i < maxLength && maxLength != 1 {
			z01.PrintRune('0')
		}
		printNextDig(i)
		start = true
	}
	z01.PrintRune('\n')
}
