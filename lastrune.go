package piscine

func LastRune(s string) rune {
	runeString := []rune(s)

	return runeString[len(runeString)-1]
}
