package main

import (
	"fmt"
	"os"
)

func isCorrectlyBracketed(s string) bool {
	stack := []rune{}

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	// ВАЖНО: по ожидаемому выводу тестов — печатаем пустую строку
	if len(os.Args) < 2 {
		fmt.Println()
		return
	}

	for _, arg := range os.Args[1:] {
		if isCorrectlyBracketed(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
