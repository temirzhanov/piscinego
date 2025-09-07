package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)

	// --- Validation: letters only, no two uppers in a row, last char not upper
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		isLower := r >= 'a' && r <= 'z'
		isUpper := r >= 'A' && r <= 'Z'

		// Only ASCII letters allowed
		if !(isLower || isUpper) {
			return s
		}
		// No two consecutive uppercase letters
		if i > 0 {
			prev := runes[i-1]
			if isUpper && (prev >= 'A' && prev <= 'Z') {
				return s
			}
		}
	}
	// Must not end with uppercase
	last := runes[len(runes)-1]
	if last >= 'A' && last <= 'Z' {
		return s
	}

	// --- Convert: insert underscores before uppercase, preserve case
	out := make([]rune, 0, len(runes)+len(runes)/2)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			if i != 0 {
				out = append(out, '_')
			}
		}
		out = append(out, r)
	}

	return string(out)
}
