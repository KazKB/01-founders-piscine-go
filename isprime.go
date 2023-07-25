package piscine

func IsPrime(nb int) bool {
	isPrime := false
	counter := 0

	for i := 1; i < nb; i++ {
		if nb%i == 0 {
			counter++
		}

		if counter > 1 {
			isPrime = false
			break
		} else if counter == 1 {
			isPrime = true
		} else if nb == 1 {
			isPrime = true
			break
		}
	}

	return isPrime
}
