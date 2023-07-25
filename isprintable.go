package piscine

func IsPrintable(s string) bool {
	isPrintable := false
	counter := 0

	for _, char := range s {
		if char > 31 && char < 127 {
			continue
		} else {
			counter++
			break
		}
	}

	if counter == 0 {
		isPrintable = true
	} else {
		isPrintable = false
	}

	return isPrintable
}
