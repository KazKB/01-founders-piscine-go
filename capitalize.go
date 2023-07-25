package piscine

func Capitalize(s string) string {
	chars := []rune(s)

	for index, char := range s {
		if index == 0 || (chars[index-1] > 31 && chars[index-1] < 48) || (chars[index-1] > 57 && chars[index-1] < 65) || (chars[index-1] > 90 && chars[index-1] < 97) || (chars[index-1] > 122 && chars[index-1] < 127) {
			if char > 96 && char < 123 {
				chars[index] = char - 32
			}
		} else if char > 64 && char < 91 {
			chars[index] = char + 32
		}
	}

	return string(chars)
}
