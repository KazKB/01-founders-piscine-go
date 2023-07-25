package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var sorted bool
	var pos, equal, neg int

	for i := 1; i < len(a); i++ {
		if compare(a[i-1], a[i]) < 0 {
			pos++
		} else if compare(a[i-1], a[i]) == 0 {
			equal++
		} else if compare(a[i-1], a[i]) > 0 {
			neg++
		}
	}

	if neg+equal == len(a)-1 || neg == len(a)-1 || pos+equal == len(a)-1 || pos == len(a)-1 {
		sorted = true
	} else {
		sorted = false
	}
	return sorted
}

func compare(a, b int) int {
	var sorted int
	if a > b {
		sorted = 1
	} else if a == b {
		sorted = 0
	} else if a < b {
		sorted = -1
	}
	return sorted
}
