package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	expr := os.Args[1]

	tokens := splitSpaces(expr)
	stack := make([]int, 0, len(tokens))

	for _, t := range tokens {
		if isOp(t) {
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var r int
			switch t {
			case "+":
				r = a + b
			case "-":
				r = a - b
			case "*":
				r = a * b
			case "/":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				r = a / b
			case "%":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				r = a % b
			}
			stack = append(stack, r)
			continue
		}

		val, ok := parseInt(t)
		if !ok {
			fmt.Println("Error")
			return
		}
		stack = append(stack, val)
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(stack[0])
}

func isOp(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "%"
}

func splitSpaces(s string) []string {
	out := []string{}
	start := -1
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && start == -1 {
			start = i
		}
		if s[i] == ' ' && start != -1 {
			out = append(out, s[start:i])
			start = -1
		}
	}
	if start != -1 {
		out = append(out, s[start:])
	}
	return out
}

func parseInt(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	sign := 1
	i := 0
	if s[0] == '+' {
		i++
	} else if s[0] == '-' {
		sign = -1
		i++
	}
	if i >= len(s) {
		return 0, false
	}
	n := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return sign * n, true
}
