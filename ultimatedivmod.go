package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a
	mod := *b

	*a = div / mod
	*b = div % mod
}
