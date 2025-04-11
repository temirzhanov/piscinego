package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	length := lenArr(arguments)

	if length == 1 {

		/*
			s := ""
			b := make([]byte, 1)
			count := 0
			for os.Stdin.Read(b); b[0] != '\n'; os.Stdin.Read(b) {
				s += string(b)
				count++
				if count > 100 {
					break
				}
			}
			if s == "" || s == "\n" {
				return
			}
			for _, j := range s {
				z01.PrintRune(j)
			}
			return
		*/

		reader := io.TeeReader(os.Stdin, os.Stdout)

		// Read all data from the TeeReader
		_, err := io.ReadAll(reader)
		if err != nil {
			return
		}
	}

	for _, arg := range arguments[1:] {
		file, err := os.Open(arg)
		if err != nil {
			errorStr := "ERROR: " + err.Error() + "\n"
			printStr(errorStr)
			os.Exit(1)
		}

		f, err := os.ReadFile(arg)

		printStr(string(f))

		if err != nil {
			printStr(err.Error() + "\n")
		}

		file.Close()
	}
}

func lenArr(arr []string) int {
	count := 0
	for range arr {
		count++
	}
	return count
}

func printStr(str string) {
	for _, letter := range str {
		z01.PrintRune(letter)
	}
}
