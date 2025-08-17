package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	pat := os.Args[1]
	text := os.Args[2]
	if len(text) == 0 {
		return
	}

	alts, ok := parsePattern(pat)
	if !ok || len(alts) == 0 {
		return
	}

	words := splitSpaces(text)
	n := 0
	for _, raw := range words {
		w := trimNonLetters(raw)
		if len(w) == 0 {
			continue
		}
		c := countDistinctContains(w, alts)
		for i := 0; i < c; i++ {
			n++
			fmt.Printf("%d: %s\n", n, w)
		}
	}
}

func parsePattern(p string) ([]string, bool) {
	if len(p) < 2 {
		return nil, false
	}
	open, close := p[0], p[len(p)-1]
	if !((open == '(' && close == ')') || (open == '[' && close == ']')) {
		return nil, false
	}
	inner := p[1 : len(p)-1]
	if len(inner) == 0 {
		return nil, false
	}
	parts := splitByBar(inner)
	for _, s := range parts {
		if len(s) == 0 {
			return nil, false
		}
	}
	return parts, true
}

func countDistinctContains(word string, alts []string) int {
	cnt := 0
	for _, a := range alts {
		if contains(word, a) {
			cnt++
		}
	}
	return cnt
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

func splitByBar(s string) []string {
	out := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	out = append(out, s[start:])
	return out
}

func trimNonLetters(s string) string {
	l, r := 0, len(s)-1
	for l <= r && !isLetter(s[l]) {
		l++
	}
	for r >= l && !isLetter(s[r]) {
		r--
	}
	if l > r {
		return ""
	}
	return s[l : r+1]
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func contains(hay, needle string) bool {
	if len(needle) == 0 {
		return true
	}
	if len(needle) > len(hay) {
		return false
	}
	for i := 0; i+len(needle) <= len(hay); i++ {
		if equals(hay[i:i+len(needle)], needle) {
			return true
		}
	}
	return false
}

func equals(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
