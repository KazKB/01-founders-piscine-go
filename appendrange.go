package piscine

func AppendRange(min, max int) []int {
	var array []int
	if min < max {
		for i := min; i < max; i++ {
			array = append(array, i)
		}
	}
	return array
}
