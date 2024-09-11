package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	sln := []rune(base)
	slb := []rune{}
	if baseChecker(base) {
		if nbr == 0 {
			z01.PrintRune('0')
			return
		}
		if nbr == -9223372036854775808 {
			z01.PrintRune('-')
			z01.PrintRune('9')
			nbr = 223372036854775808
		}
		if nbr < 0 {
			z01.PrintRune('-')
			nbr *= -1
		}

		for nbr > 0 {
			slb = append(slb, sln[nbr%len(sln)])
			nbr /= len(sln)
		}
	} else {
		slb = []rune{'V', 'N'}
	}

	for i := len(slb) - 1; i >= 0; i-- {
		z01.PrintRune(slb[i])
	}
}

func baseChecker(s string) bool {
	sl := []rune(s)
	if len(sl) < 2 {
		return false
	}
	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] == sl[j] || sl[i] == '-' || sl[j] == '+' {
				return false
			}
		}
	}
	return true
}
