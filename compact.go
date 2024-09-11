package piscine

func Compact(ptr *[]string) int {
	sl := *ptr
	com := make([]string, 0, len(sl))
	c := 0
	for _, v := range sl {
		if v != "" {
			com = append(com, v)
			c++
		}
	}
	*ptr = com
	return c
}
