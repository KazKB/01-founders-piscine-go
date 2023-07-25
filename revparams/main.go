package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args

	for i := len(params) - 1; i > 0; i-- {
		runeArray := []rune(os.Args[i])

		for j := range runeArray {
			z01.PrintRune(runeArray[j])
		}
		z01.PrintRune('\n')
	}
}
