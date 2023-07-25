package piscine

func NRune(s string, n int) rune {
	runeString := []rune(s)

	if n <= len(runeString) && n > 0 {
		return runeString[n-1]
	} else {
		return 0
	}
}
