package piscine

func CountIf(f func(string) bool, tab []string) int {
	var numeric []bool

	var hasNumbers int

	for i := range tab {
		numeric = append(numeric, f(tab[i]))
	}

	for i := range numeric {
		if numeric[i] {
			hasNumbers++
		}
	}
	return hasNumbers
}
