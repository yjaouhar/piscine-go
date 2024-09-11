package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var sl []int
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	for n > 0 {
		sl = append(sl, n%10)
		n /= 10

	}
	for i := len(sl) - 1; i >= 0; i-- {
		z01.PrintRune(rune(sl[i] + '0'))
	}
}
