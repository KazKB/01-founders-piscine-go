package piscine

func Sqrt(nb int) int {
	product := 0

	for i := 1; i <= nb; i++ {
		if i*i == nb {
			product = i
			break
		} else {
			product = 0
		}
	}

	return product
}
