package piscine

func Unmatch(a []int) int {
	var sl int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				sl++
			}
		}
		if sl%2 != 0 {
			return a[i]
		}
	}
	return -1
}
