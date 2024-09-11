package piscine

func Rot14(s string) string {
	sl := []rune(s)
	var str string

	for _, v := range sl {
		if (v >= 'a' && v < 'm') || (v >= 'A' && v < 'M') {
			str += string(v + 14)
		} else if (v >= 'm' && v <= 'z') || (v >= 'M' && v <= 'Z') {
			str += string(v - 12)
		} else {
			str += string(v)
		}
	}
	return str
}
