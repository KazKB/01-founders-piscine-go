package piscine

func IsAlpha(s string) bool {
	isAlpha := false
	counter := 0

	for _, char := range s {
		if (char > 64 && char < 91) || (char > 96 && char < 123) || (char > 47 && char < 58) {
			continue
		} else {
			counter++
			break
		}
	}

	if counter == 0 {
		isAlpha = true
	} else {
		isAlpha = false
	}

	return isAlpha
}
