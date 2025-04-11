package piscine

func LoafOfBread(str string) string {
	res := ""
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "invalid Output\n"
	}
	count := 0
	for _, letter := range str {
		if letter != ' ' && count != 5 {
			res += string(letter)
			count++
		} else if count == 5 {
			res += " "
			count = 0
		}
	}
	if len(res) > 0 && res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}
	return res + "\n"
}
