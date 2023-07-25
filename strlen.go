package piscine

func StrLen(s string) int {
	var count int = 0

	string := []rune(s)

	for i := range string {
		count = i + 1
	}

	return count
}
