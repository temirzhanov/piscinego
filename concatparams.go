package piscine

func ConcatParams(args []string) string {
	str := ""
	lenArg := lenArr(args)

	for ind, word := range args {
		str += word
		if ind != lenArg-1 {
			str += "\n"
		}
	}

	return str
}
