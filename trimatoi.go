package piscine

func TrimAtoi(s string) int {
	t := 0
	multi := 1

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			t = t*10 + int(s[i]-'0')
		} else if !(s[i] >= '0' || s[i] <= '9') {
			return t
		}
		if s[i] == '-' {
			if s[i-1] >= 'a' && s[i-1] <= 'z' {
				multi = -1
			}
		}
	}

	return t * multi
}
