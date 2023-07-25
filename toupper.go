package piscine

func ToUpper(s string) string {
	upper := ""
	for _, char := range s {
		if char > 96 && char < 123 {
			upper += string(char - 32)
		} else {
			upper += string(char)
		}
	}

	return upper
}
