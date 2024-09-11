package piscine

func RecursivePower(nb int, power int) int {
	var p int = 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		p *= nb * RecursivePower(nb, power-1)
	}
	return p
}
