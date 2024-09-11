package piscine

func BasicAtoi(s string) int {
	t := []rune(s)
	a := 0
	for i := 0; i < len(t); i++ {
		if t[i] >= '0' && t[i] <= '9' {
			a = a*10 + int(t[i]-'0')
		} else {
			a = 0
			break
		}
	}

	return a
}
