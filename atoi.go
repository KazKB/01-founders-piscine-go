package piscine

func Atoi(s string) int {
	num := 0
	minus := false
	plus := false
	minusCounter := 0
	plusCounter := 0

	for index, char := range s {
		if index == 0 && char == 43 {
			plus = true
			plusCounter = 1
			continue
		} else if index == 0 && char == 45 {
			minus = true
			minusCounter = 1
			continue
		} else if char == 43 {
			plusCounter++
			continue
		} else if char == 45 {
			minusCounter++
			continue
		} else if char < 48 || char > 57 {
			num = 0
			break
		} else if char > 47 || char < 58 {
			num = num*10 + int(char-48)
		}
	}

	if minus && minusCounter == 1 {
		num = num * -1
	} else if (minusCounter > 0 && !minus) || (minusCounter > 1 && minus) {
		num = 0
	}

	if plusCounter > 1 {
		num = 0
	}

	if !plus && plusCounter > 0 {
		num = 0
	}

	if plusCounter > 0 && minusCounter > 0 {
		num = 0
	}

	return num
}
