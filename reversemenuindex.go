package piscine

func ReverseMenuIndex(menu []string) []string {
	sl := make([]string, len(menu))
	for i := range menu {
		sl[len(menu)-1-i] = menu[i]
	}
	return sl
}
