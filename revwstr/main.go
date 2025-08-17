package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}
	var arr []string
	word := ""
	check := false
	for _, char := range args[0] {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			word += string(char)
			check = true
		} else {
			check = false
			arr = append(arr, word)
			word = ""
		}
	}
	if check {
		arr = append(arr, word)
	}
	var res []string
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}

	printArray(res)
	return
}

func printArray(arr []string) {
	for i := 0; i < len(arr); i++ {
		for _, char := range arr[i] {
			z01.PrintRune(char)
		}
		if i != len(arr)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
