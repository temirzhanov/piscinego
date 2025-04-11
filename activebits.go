package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1 // Right shift n by 1
	}
	return count
}

/*
func ActiveBits(n int) int {
	return countOnes(returnNbrBase(n, "01"))
}

func countOnes(str string) int {
	count := 0
	for _, letter := range str {
		if letter == '1' {
			count++
		}
	}
	return count
}

func returnToBase(nbr int, base []rune, size int, sign int) string {
	res := ""
	if nbr < size && nbr > -1*size {
		res += string(base[nbr*sign])
		return res
	}
	res += returnToBase(nbr/size, base, size, sign)
	res += string(base[nbr%size*sign])
	return res
}

func returnNbrBase(nbr int, base string) string {
	size := 0
	isValid := true
	res := ""
	for _, val := range base {
		if val == '-' || val == '+' {
			isValid = false
			size = 0
			break
		}
		size++
	}
	if size < 2 {
		res += "NV"
		return res
	}
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if base[j] == base[i] {
				isValid = false
				break
			}
		}
	}
	if !isValid {
		res += "NV"
		return res
	}
	dot := 1
	if nbr < 0 {
		dot *= -1
		res += "-"
	}
	res += returnToBase(nbr, []rune(base), size, dot)
	return res
}
*/
