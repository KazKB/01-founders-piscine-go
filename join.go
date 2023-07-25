package piscine

func Join(elems []string, sep string) string {
	concat := ""

	for index := range elems {
		if index < len(elems)-1 {
			concat += elems[index] + sep
		} else {
			concat += elems[index]
		}
	}

	return concat
}
