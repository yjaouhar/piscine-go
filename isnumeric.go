package piscine

func IsNumeric(s string) bool {
	var a bool
	for _, v := range s {
		if v >= '0' && v <= '9' {
			a = true
		} else {
			a = false
			break
		}
	}
	return a
}
