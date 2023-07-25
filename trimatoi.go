package piscine

func TrimAtoi(s string) int {
	num := 0
	minus := false

	for index, char := range s {
		if char == 45 && num == 0 {
			minus = true
		} else if char < 48 || char > 57 {
			index++
		} else if char > 47 || char < 58 {
			num = num*10 + int(char-48)
		}
	}

	if minus {
		num = -num
	}

	return num
}
