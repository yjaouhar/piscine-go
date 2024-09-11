package piscine

func Split(s, sep string) []string {
	sl := []string{}
	t := ""
	sepLen := len(sep)

	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			sl = append(sl, t)
			t = ""
			i += sepLen - 1
		} else {
			t += string(s[i])
		}
	}

	if t != "" {
		sl = append(sl, t)
	}

	return sl
}
