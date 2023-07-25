package piscine

func Index(s string, toFind string) int {
	runeString := []rune(s)
	findString := []rune(toFind)
	index := 0
	counter := 0

	if len(runeString) > 0 && len(findString) > 0 {
		for i := 0; i < len(runeString); i++ {
			if runeString[i] == findString[0] {
				k := i
				for j := 0; j < len(findString); j++ {
					if runeString[k] == findString[j] {
						counter++
						k++
					}
				}

				if counter == len(findString) {
					index = i
					break
				}
			} else {
				index = -1
			}
		}
	} else if len(findString) == 0 {
		index = 0
	} else {
		index = -1
	}

	return index
}
