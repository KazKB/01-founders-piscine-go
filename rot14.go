package piscine

func Rot14(s string) string {
	var excess, spacesLeft int
	var encrypted []string

	for i := 0; i < len(s); i++ {
		if s[i] > 76 && s[i] < 91 {
			spacesLeft = 91 - int(s[i])
			excess = 65 + (14 - spacesLeft)
			encrypted = append(encrypted, string(byte(excess)))
		} else if s[i] > 108 && s[i] < 123 {
			spacesLeft = 123 - int(s[i])
			excess = 97 + (14 - spacesLeft)
			encrypted = append(encrypted, string(byte(excess)))
		} else if (s[i] > 64 && s[i] < 77) || (s[i] > 96 && s[i] < 109) {
			encrypted = append(encrypted, string(s[i]+14))
		} else {
			encrypted = append(encrypted, string(s[i]))
		}
	}

	s = ""

	for i := range encrypted {
		s += encrypted[i]
	}

	return s
}
