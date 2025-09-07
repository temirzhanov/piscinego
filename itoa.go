package piscine

func Itoa(n int) string {
	res := ""
	sign := ""
	if n < 0 {
		sign += "-"
		n = -n
	} else if n == 0 {
		res += "0"
	}

	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}

	res = sign + res

	return res

}
