package piscine

func SplitWhiteSpaces(s string) []string {
	sl := []string{}
	t := ""

	for i, v := range s {
		if !(v == ' ' || v == '\n' || v == '\t') {
			t += string(v)
		} else {
			if t != "" {
				sl = append(sl, t)
				t = ""
			}
		}

		if i == len(s)-1 && t != "" {
			sl = append(sl, t)
			t = ""
		}
	}

	return sl
}
