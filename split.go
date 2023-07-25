package piscine

func Split(s, sep string) []string {
	var words []string

	char := []rune(s)
	sepChar := []rune(sep)

	var word string
	var counter, j int

	for i := 0; i < len(char); i++ {
		if char[i] == sepChar[0] {
			j = i
			counter = 0

			for k := 0; k < len(sepChar); k++ {
				if char[j] == sepChar[k] {
					counter++
					j++
				}
			}

			if counter == len(sepChar) {
				i += counter - 1
				if word != "" {
					words = append(words, word)
				}
				word = ""
			} else {
				word += string(char[i])
			}
		} else if i == len(char)-1 {
			word += string(char[i])
			if word != "" {
				words = append(words, word)
			}
		} else {
			word += string(char[i])
		}
	}
	return words
}
