package piscine

func ToUpper(s string) string {
	a := []byte(s)
	for i := 0; i < len(a); i++ {
		if a[i] >= 'a' && a[i] <= 'z' {
			a[i] -= 32
		}
	}
	return string(a)
}
