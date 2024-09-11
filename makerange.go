package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	b := max - min
	sl := make([]int, b)
	for i := 0; i < b; i++ {
		sl[i] = i + min
	}
	return sl
}
