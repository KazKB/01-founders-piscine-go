package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		word := []rune(a[i])
		for j := 0; j < len(word); j++ {
			z01.PrintRune(word[j])
		}
		z01.PrintRune('\n')
	}
}
