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
		w := trimWord(raw)
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

// trimWord:
// - remove leading non-letters
// - keep from first letter through the last letter
// - keep exactly ONE trailing digit *only if* there are NO digits inside the [first..last letter] span
func trimWord(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Find first letter (left bound)
	l := 0
	for l < len(s) && !isLetter(s[l]) {
		l++
	}
	if l >= len(s) {
		return ""
	}

	// Find last letter (right bound)
	lastLetter := -1
	for r := len(s) - 1; r >= l; r-- {
		if isLetter(s[r]) {
			lastLetter = r
			break
		}
	}
	if lastLetter == -1 {
		return ""
	}

	// Check if there is any digit within the core [l..lastLetter]
	hasDigitInCore := false
	for i := l; i <= lastLetter; i++ {
		if isDigit(s[i]) {
			hasDigitInCore = true
			break
		}
	}

	end := lastLetter + 1
	// Keep one trailing digit only if core has no digits
	if !hasDigitInCore && end < len(s) && isDigit(s[end]) {
		end++
	}

	return s[l:end]
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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
