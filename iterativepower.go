package piscine

func IterativePower(nb int, power int) int {
	var p int = 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for i := 1; i <= power; i++ {
			p *= nb
		}
	}
	return p
}
