package piscine

func Unmatch(a []int) int {
	var unmatched int = -1
	var counter int

	for i := 0; i < len(a); i++ {
		counter = 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				counter++
			}
		}
		if counter%2 == 1 {
			unmatched = a[i]
			break
		}
	}

	return unmatched
}
