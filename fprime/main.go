package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	first := true

	for div := 2; div*div <= num; div++ {
		for num%div == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(div)
			first = false
			num /= div
		}
	}

	if num > 1 {
		if !first {
			fmt.Print("*")
		}
		fmt.Print(num)
	}

	fmt.Println()
}
