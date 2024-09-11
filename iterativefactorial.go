package piscine

func IterativeFactorial(nb int) int {
	s := 1
	for i := 1; i <= nb; i++ {
		s = s * i
	}
	return s
}
