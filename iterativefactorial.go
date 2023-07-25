package piscine

func IterativeFactorial(nb int) int {
	factorial := 1

	if nb >= 0 && nb <= 2147483647 {
		for i := 1; i <= nb; i++ {
			factorial *= i
		}

		return factorial
	} else {
		return 0
	}
}
