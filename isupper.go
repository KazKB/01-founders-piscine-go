package piscine

func IsUpper(s string) bool {
	isUpper := false
	counter := 0

	for _, char := range s {
		if char > 64 && char < 91 {
			continue
		} else {
			counter++
			break
		}
	}

	if counter == 0 {
		isUpper = true
	} else {
		isUpper = false
	}

	return isUpper
}
