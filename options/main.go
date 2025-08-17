package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	
	for _, a := range args {
		if a == "" {
			fmt.Println("Invalid Option")
			return
		}
		if len(a) >= 2 && a[0] == '-' && a[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
	}

	var mask uint32 = 0

	for _, a := range args {
		if len(a) < 2 || a[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		for i := 1; i < len(a); i++ {
			c := a[i]
			if c < 'a' || c > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			shift := uint32(c - 'a')
			mask |= 1 << shift
		}
	}

	printBits(mask)
}

func printBits(v uint32) {
	for i := 31; i >= 0; i-- {
		if (v & (1 << uint(i))) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
