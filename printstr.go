package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	string := []byte(s)

	for _, char := range string {
		z01.PrintRune(rune(char))
	}
}
