package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var s []int
	var a int
	if n >= 0 {
		if n == 0 {
			z01.PrintRune('0')
		}
		for i := 0; i <= 19; i++ {
			if n == 0 {
				break
			}
			a = n % 10
			n = n / 10
			s = append(s, a)

		}
		for j := 0; j < len(s)-1; j++ {
			for k := 0; k < len(s)-1-j; k++ {
				if s[k] > s[k+1] {
					s[k], s[k+1] = s[k+1], s[k]
				}
			}
		}
		for _, l := range s {
			z01.PrintRune(rune(l + '0'))
		}
	}
}
