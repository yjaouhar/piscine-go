package piscine

func StrRev(s string) string {
	t := ""
	a := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		t += string(a[i])
	}
	return t
}
