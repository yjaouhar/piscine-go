package piscine

func IsPrintable(s string) bool {
	var a bool
	for _, v := range s {
		if v >= '!' && v <= '~' {
			a = true
		} else {
			a = false
			break
		}
	}
	return a
}
