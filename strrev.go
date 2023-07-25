package piscine

func StrRev(s string) string {
	forward := []rune(s)
	reverse := ""

	for i := len(forward) - 1; i >= 0; i-- {
		reverse += string(forward[i])
	}

	return reverse
}
