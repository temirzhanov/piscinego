package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	srcRaw := os.Args[1]

	// Keep only valid Brainfuck instructions (ignore ',' and comments)
	code := make([]rune, 0, len(srcRaw))
	for _, c := range srcRaw {
		if c == '>' || c == '<' || c == '+' || c == '-' ||
			c == '.' || c == '[' || c == ']' {
			code = append(code, c)
		}
	}

	// Build bracket-matching map
	match := make(map[int]int)
	stack := []int{}
	for i, c := range code {
		if c == '[' {
			stack = append(stack, i)
		} else if c == ']' {
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			match[open] = i
			match[i] = open
		}
	}
	// valid program => stack empty

	// Tape
	const tapeSize = 2048
	tape := make([]byte, tapeSize)
	ptr := 0
	pc := 0

	for pc < len(code) {
		switch code[pc] {
		case '>':
			ptr++
			if ptr >= tapeSize {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = tapeSize - 1
			}
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			z01.PrintRune(rune(tape[ptr]))
		case '[':
			if tape[ptr] == 0 {
				pc = match[pc]
			}
		case ']':
			if tape[ptr] != 0 {
				pc = match[pc]
			}
		}
		pc++
	}
}
