package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	var EvenMsg string = "I have an even number of arguments"
	var OddMsg string = "I have an odd number of arguments"
	var arg []string = os.Args

	if isEven(len(arg)) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
