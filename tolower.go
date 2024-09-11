package piscine

func ToLower(s string) string {
	a := []byte(s)
	for i := 0; i < len(a); i++ {
		if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] += 32
		}
	}
	return string(a)
}
