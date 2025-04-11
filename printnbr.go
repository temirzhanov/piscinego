package piscine

import (
	"github.com/01-edu/z01"
)

func CheckNum(numb int) {
	Rnum := '0'
	if numb == 0 {
		z01.PrintRune(Rnum)
		return
	}
	for i := 1; i <= numb%10; i++ {
		Rnum++
	}
	for i := -1; i >= numb%10; i-- {
		Rnum++
	}

	if numb/10 != 0 {
		CheckNum(numb / 10)
	}
	z01.PrintRune(Rnum)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	CheckNum(n)
}
