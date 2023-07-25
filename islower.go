package piscine

func IsLower(s string) bool {
	isLower := false
	counter := 0

	for _, char := range s {
		if char > 96 && char < 123 {
			continue
		} else {
			counter++
			break
		}
	}

	if counter == 0 {
		isLower = true
	} else {
		isLower = false
	}

	return isLower
}
