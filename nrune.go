package piscine

func NRune(s string, n int) rune {
	v := []rune(s)

	if n < 0 || n > len(s) {
		return 0
	} else {
		return v[n-1]
	}
}
