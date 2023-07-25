package piscine

func ActiveBits(n int) int {
	var bits int
	for n > 0 {
		if n%2 == 1 {
			bits++
		}
		n /= 2
	}

	return bits
}
