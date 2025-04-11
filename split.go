package piscine

func Split(str, charset string) []string {
	char_len := StrLen(charset)
	str_len := StrLen(str)
	size := 0
	for ind := 0; ind <= str_len-char_len; ind++ {
		if string(str[ind:ind+char_len]) == charset {
			size++
		}
	}
	resArr := make([]string, size+1)
	i := 0
	start := 0
	ind := 0
	for ; ind <= str_len-char_len; ind++ {
		if string(str[ind:ind+char_len]) == charset {
			resArr[i] = string(str[start:ind])
			i++
			start = ind + char_len
		}
		if ind == str_len-char_len {
			resArr[i] = string(str[start:])
		}
	}
	return resArr
}
