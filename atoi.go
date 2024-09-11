package piscine

func Atoi(s string) int {
	sl := 0
	multi := 1

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			sl = sl*10 + int(s[i]-'0')
		} else if s[i] == '-' && i == 0 {
			multi = -1
		} else if s[i] == '+' && i == 0 {
			continue
		} else {
			multi = 0
		}
	}
	return sl * multi
}
