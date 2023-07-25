package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	runeArray := []rune(name)

	for i := 2; i < len(runeArray); i++ {
		z01.PrintRune(runeArray[i])
	}
	z01.PrintRune('\n')
}
