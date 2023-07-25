package piscine

func ToLower(s string) string {
	lower := ""
	for _, char := range s {
		if char > 64 && char < 91 {
			lower += string(char + 32)
		} else {
			lower += string(char)
		}
	}

	return lower
}
