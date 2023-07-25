package piscine

func CollatzCountdown(start int) int {
	var steps, n int

	n = start

	if start <= 0 {
		steps = -1
	} else {
		for n > 1 {
			if n%2 == 0 {
				n /= 2
				steps++
			} else if n%2 == 1 {
				n = 3*n + 1
				steps++
			}
		}
	}

	return steps
}
