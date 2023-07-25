package piscine

func Any(f func(string) bool, a []string) bool {
	var numeric []bool

	var boolean bool

	for i := range a {
		numeric = append(numeric, f(a[i]))
	}

	for i := range numeric {
		if numeric[i] {
			boolean = true
		}
	}
	return boolean
}
