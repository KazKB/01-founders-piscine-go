package piscine

func Map(f func(int) bool, a []int) []bool {
	var prime []bool

	for i := range a {
		prime = append(prime, f(a[i]))
	}
	return prime
}
