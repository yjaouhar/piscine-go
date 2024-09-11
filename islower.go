package piscine

func IsLower(s string) bool {
	var a bool
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			a = true
		} else {
			a = false
			break
		}
	}
	return a
}
