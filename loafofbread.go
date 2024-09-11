package piscine

func LoafOfBread(str string) string {
	var sl []rune
	s := str
	if s == "" {
		sl = append(sl, '\n')
		return string(sl)
	}

	if len(s) < 5 {
		sl = append(sl, []rune("Invalid Output\n")...)
		return string(sl)
	}

	j := 0
	for i := 0; i < len(s); i++ {
		if j < 5 && rune(s[i]) == ' ' {
			continue
		}
		if j == 5 {
			if i != len(s)-1 && s[i+1] == ' ' {
				continue
			}
			if i == len(s)-1 {
				break
			}
			sl = append(sl, ' ')
			j = 0
			continue
		}
		sl = append(sl, rune(s[i]))
		j++
	}

	sl = append(sl, '\n')
	return string(sl)
}
