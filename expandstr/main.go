package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	check := false
	str := ""

	for i := 0; i < len(os.Args[1]); i++ {
		if os.Args[1][i] != ' ' {
			if check && str != "" {
				str += "   "
			}
			str += string(os.Args[1][i])
			check = false
		} else {
			check = true
		}
	}

	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
