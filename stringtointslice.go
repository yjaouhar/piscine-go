package piscine

func StringToIntSlice(str string) []int {
	var sl []int
	for _, v := range str {
		sl = append(sl, int(v))
	}
	return sl
}
