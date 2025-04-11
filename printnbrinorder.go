package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}

	for i := 0; i <= 9; i++ {
		res := n
		for res > 0 {
			if res%10 == i {
				z01.PrintRune(rune(res%10 + 48))
			}
			res /= 10
		}
	}
}
