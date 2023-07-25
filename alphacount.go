package piscine

func AlphaCount(s string) int {
	runeString := []rune(s)
	count := 0

	for i := 0; i < len(runeString); i++ {
		if (runeString[i] > 64 && runeString[i] < 91) || (runeString[i] > 96 && runeString[i] < 123) {
			count++
		}
	}

	return count
}
