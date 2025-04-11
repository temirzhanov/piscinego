package piscine

func ShoppingListSort(slice []string) []string {
	for ind := 0; ind < len(slice)-1; ind++ {
		for jnd := ind + 1; jnd < len(slice); jnd++ {
			if len(slice[ind]) > len(slice[jnd]) {
				slice[ind], slice[jnd] = slice[jnd], slice[ind]
			}
		}
	}

	return slice
}
