package piscine

func Itoa(n int) string {
	var digit []int
	var temp int = n
	var number string

	if n < 0 {
		number += "-"
		n *= -1
	}

	for n > 0 {
		temp = n % 10
		digit = append(digit, temp)
		n /= 10
	}

	for i := len(digit) - 1; i >= 0; i-- {
		number = number + string(rune(digit[i]+48))
	}

	return number
}
