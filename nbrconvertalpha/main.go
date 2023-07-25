package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	num := 0

	for _, char := range s {
		if char < 48 || char > 57 {
			num = 0
			break
		} else {
			num = num*10 + int(char-48)
		}
	}

	return num
}

func main() {
	params := os.Args

	num := ""

	if len(params) < 2 {
	} else if params[1] == "--upper" {
		for i := 2; i < len(params); i++ {
			num = params[i]

			if BasicAtoi(num) >= 1 && BasicAtoi(num) <= 26 {
				z01.PrintRune(rune(BasicAtoi(num) + 64))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	} else {
		for i := 1; i < len(params); i++ {
			num = params[i]

			if BasicAtoi(num) >= 1 && BasicAtoi(num) <= 26 {
				z01.PrintRune(rune(BasicAtoi(num) + 96))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
