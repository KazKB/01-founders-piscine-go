package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	char := []rune(s)
	var word string

	for i := 0; i < len(char); i++ {
		if char[i] == 32 || char[i] == 9 || char[i] == '\n' {
			if word != "" {
				words = append(words, word)
			}
			word = ""
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
