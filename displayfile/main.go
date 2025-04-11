package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	lenArg := lenArr(args)
	if lenArg == 1 {
		fmt.Println("File name missing")
	} else if lenArg > 2 {
		fmt.Println("Too many arguments")
	} else {
		content, err := os.ReadFile(args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(content))

	}
}

func lenArr(arr []string) int {
	count := 0
	for range arr {
		count++
	}
	return count
}
