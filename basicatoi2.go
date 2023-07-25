package piscine

func BasicAtoi2(s string) int {
	num := 0

	for _, char := range s {
		if char < 48 || char > 57 {
			num = 0
			break
		} else {
			num = num*10 + int(char-48)
		}
	}

	return num
}
