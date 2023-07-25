package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := i; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					if ((i < k) || (i == k && j < l)) && (i != 9 || j != 9) {
						z01.PrintRune(rune(i + 48))
						z01.PrintRune(rune(j + 48))
						z01.PrintRune(' ')
						z01.PrintRune(rune(k + 48))
						z01.PrintRune(rune(l + 48))
						if i == 9 && j == 8 && k == 9 && l == 9 {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
