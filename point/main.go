package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	var pointX, pointY []rune
	var tempX, tempY int

	for i := 0; i <= 1; i++ {
		tempX = points.x % 10
		tempY = points.y % 10

		pointX = append(pointX, rune(tempX)+48)
		pointY = append(pointY, rune(tempY)+48)

		points.x = points.x / 10
		points.y = points.y / 10
	}

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(pointX[1])
	z01.PrintRune(pointX[0])
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(pointY[1])
	z01.PrintRune(pointY[0])
	z01.PrintRune('\n')
}
