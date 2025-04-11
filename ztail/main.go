package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		os.Exit(1)
	}

	offset, err := Atoi(args[1])
	if err != nil || offset <= 0 {
		fmt.Printf("tail: invalid number of bytes: '%s'\n", args[1])
		os.Exit(1)
	}

	files := args[2:]
	multiple := len(files) > 1
	exitCode := 0

	for i, file := range files {
		content, err := ReadLastBytes(file, offset)
		if err != nil {
			// IMPORTANT: print raw error, as required by test case
			fmt.Println(err)
			exitCode = 1
			continue
		}

		if multiple {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", file)
		}
		fmt.Print(content)
	}

	os.Exit(exitCode)
}

func ReadLastBytes(filename string, count int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	size := info.Size()
	start := size - int64(count)
	if start < 0 {
		start = 0
		count = int(size)
	}

	buffer := make([]byte, count)
	_, err = file.ReadAt(buffer, start)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func Atoi(s string) (int, error) {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid number")
		}
		n = n*10 + int(c-'0')
	}
	return n, nil
}
