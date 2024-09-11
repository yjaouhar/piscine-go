package piscine

func IsUpper(s string) bool {
	var a bool
	for _, v := range s {
		if v >= 'A' && v < 'Z' {
			a = true
		} else {
			a = false
		}
	}
	return a
}
