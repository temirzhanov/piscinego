package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil || num <= 1 {
		return
	}

	// factorization
	div := 2
	first := true
	for num > 1 {
		if num%div == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(div)
			num /= div
			first = false
		} else {
			div++
		}
	}
	fmt.Println()
}
