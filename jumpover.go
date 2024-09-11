package piscine

func JumpOver(str string) string {
	sl := []rune(str)
	var stok []rune
	if len(sl) == 0 || len(sl) < 3 {
		return "\n"
	}
	for i := 2; i < len(sl); i++ {
		stok = append(stok, sl[i])
		i += 2
	}
	stok = append(stok, '\n')
	return string(stok)
}
