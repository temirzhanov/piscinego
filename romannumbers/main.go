package main

import (
	"fmt"
	"os"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func parseInt(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return n, true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("\n")
		return
	}

	raw := os.Args[1]
	n, ok := parseInt(raw)
	if !ok || n == 0 || n >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}

	vals := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4,
		1,
	}
	rom := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I",
	}

	calc := []string{
		"M", "(M-C)", "D", "(D-C)",
		"C", "(C-X)", "L", "(L-X)",
		"X", "(X-I)", "V", "(V-I)",
		"I",
	}

	remain := n
	calcTokens := []string{}
	romanOut := ""

	for i := 0; i < len(vals); i++ {
		for remain >= vals[i] {
			remain -= vals[i]
			calcTokens = append(calcTokens, calc[i])
			romanOut += rom[i]
		}
	}

	for i := 0; i < len(calcTokens); i++ {
		if i > 0 {
			fmt.Printf("+")
		}
		fmt.Printf("%s", calcTokens[i])
	}
	fmt.Printf("\n")

	fmt.Printf("%s\n", romanOut)
}
