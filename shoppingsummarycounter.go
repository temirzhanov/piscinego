package piscine

func ShoppingSummaryCounter(summary string) map[string]int {
	itemCounts := make(map[string]int)
	if summary == "" {
		itemCounts[""] = 1
		return itemCounts
	}

	item := ""

	for i := 0; i < len(summary); i++ {
		char := summary[i]
		if char == ' ' {
			if item == "" {
				itemCounts[""]++
			}
			if item != "" {
				itemCounts[item]++
			}
			item = ""
		} else {
			item += string(char)
		}
	}

	if item != "" {
		itemCounts[item]++
	}

	if item == "" && len(summary) > 0 && summary[len(summary)-1] == ' ' {
		itemCounts[""]++
	}

	return itemCounts
}
