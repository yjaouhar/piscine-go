package piscine

func IsAlpha(s string) bool {
	var a bool
	for _, v := range s {
		if v >= 'A' && v <= 'Z' || v >= 'a' && v <= 'z' || v >= 0 && v <= 9 {
			a = true
		} else {
			a = false
			break
		}
	}
	return a
}
