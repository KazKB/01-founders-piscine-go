package piscine

func MakeRange(min, max int) []int {
	var array []int

	if min < max {
		array, j := make([]int, max-min), 0
		for i := min; i < max; i++ {
			array[j] = i
			j++
		}
		return array
	}
	return array
}
