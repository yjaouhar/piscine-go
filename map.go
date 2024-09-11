package piscine

func Map(f func(int) bool, a []int) []bool {
	sl := []bool{}
	for i := 0; i < len(a); i++ {
		sl = append(sl, f(a[i]))
	}
	return sl
}
