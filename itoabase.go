package piscine

func ItoaBase(value, base int) string {
	const chars = "0123456789ABCDEF"

	if base < 2 || base > 16 {
		return ""
	}
	if value == 0 {
		return "0"
	}

	neg := value < 0

	// Work with unsigned magnitude to avoid overflow on MinInt
	var u uint64
	if neg {
		u = uint64(-(value + 1))
		u++
	} else {
		u = uint64(value)
	}

	// Build digits in reverse order
	var buf []byte
	for u > 0 {
		d := u % uint64(base)
		buf = append(buf, chars[d])
		u /= uint64(base)
	}

	// Reverse in place
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	if neg {
		return "-" + string(buf)
	}
	return string(buf)
}
