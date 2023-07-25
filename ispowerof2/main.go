package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	var args []string = os.Args

	if len(args) == 2 {
		var power string

		num, _ := strconv.Atoi(args[1])

		if num > 1 || num < 0 {
			for num > 1 {
				if num%2 == 0 {
					power = "true"
					num /= 2
				} else {
					power = "false"
					break
				}
			}
		} else if num == 1 {
			power = "true"
		}

		for i := range power {
			z01.PrintRune(rune(power[i]))
		}
		z01.PrintRune('\n')
	}
}
