package piscine

func IterativePower(nb int, power int) int {
	product := 1

	if power > 0 {
		for i := 1; i <= power; i++ {
			product *= nb
		}
		return product
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}
