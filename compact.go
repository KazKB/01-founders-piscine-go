package piscine

func Compact(ptr *[]string) int {
	strings := []string(*ptr)

	var nonZero []string
	var numNonZero int

	for i := range *ptr {
		if strings[i] != "" {
			nonZero = append(nonZero, strings[i])
			numNonZero++
		} else {
			continue
		}
	}

	*ptr = nonZero

	return numNonZero
}
