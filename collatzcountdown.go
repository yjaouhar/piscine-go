package piscine

func CollatzCountdown(start int) int {
	var c int
	if start <= 0 {
		return -1
	} else {
		c = 0
		for start != 1 {
			if start%2 == 0 {
				c++
				start /= 2

			} else if start%2 != 0 {
				c++
				start = start*3 + 1
			}
		}
	}

	return c
}
