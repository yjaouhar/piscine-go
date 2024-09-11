package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for _, i := range tab {
		if f(i) {
			count++
		}
	}
	return count
}
