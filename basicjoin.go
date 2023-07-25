package piscine

func BasicJoin(elems []string) string {
	concat := ""

	for index := range elems {
		concat += elems[index]
	}

	return concat
}
