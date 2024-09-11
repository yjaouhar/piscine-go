package piscine

func Abort(a, b, c, d, e int) int {
	var sl []int = []int{a, b, c, d, e}
	for i := 0; i < len(sl)-1; i++ {
		for j := 0; j < len(sl)-i-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
	return sl[2]
}
