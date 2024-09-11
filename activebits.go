package piscine

func ActiveBits(n int) int {
	m := 0
	if n < 0 {
		n *= -1
	}
	for n != 0 {
		if n%2 == 1 {
			m++
		}
		n /= 2
	}
	return m
}
