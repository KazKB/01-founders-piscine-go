package piscine

func IsNumeric(s string) bool {
	isNumeric := false
	counter := 0

	for _, char := range s {
		if char > 47 && char < 58 {
			continue
		} else {
			counter++
			break
		}
	}

	if counter == 0 {
		isNumeric = true
	} else {
		isNumeric = false
	}

	return isNumeric
}
