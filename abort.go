package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
		}
	}

	return numbers[2]
}
