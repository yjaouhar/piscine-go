package piscine

func AlphaCount(s string) int {
	r := []rune(s)
	var index int
	for _, v := range r {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			index++
		}
	}

	return index
}
