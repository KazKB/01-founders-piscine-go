package piscine

func Max(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a); j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}

	return a[len(a)-1]
}
