package piscine

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	dot := -1
	hasDigit := false

	for i := 0; i < len(dec); i++ {
		c := dec[i]

		if c == '-' {
			if i != 0 || len(dec) == 1 {
				return dec + "\n"
			}
			continue
		}

		if c >= '0' && c <= '9' {
			hasDigit = true
			continue
		}

		if c == '.' && dot == -1 {
			dot = i
			continue
		}

		return dec + "\n"
	}

	if !hasDigit {
		return dec + "\n"
	}

	// Нет точки
	if dot == -1 {
		return dec + "\n"
	}

	// Проверка дробной части
	if dot == len(dec)-1 {
		return trimZeros(dec[:dot]) + "\n"
	}

	onlyZero := true
	for i := dot + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			onlyZero = false
			break
		}
	}
	if onlyZero {
		return trimZeros(dec[:dot]) + "\n"
	}

	// Убираем точку
	res := ""
	for i := 0; i < len(dec); i++ {
		if dec[i] != '.' {
			res += string(dec[i])
		}
	}

	return trimZeros(res) + "\n"
}

func trimZeros(s string) string {
	sign := ""
	i := 0

	if s[0] == '-' {
		sign = "-"
		i = 1
	}

	for i < len(s) && s[i] == '0' {
		i++
	}

	if i == len(s) {
		return "0"
	}

	return sign + s[i:]
}
