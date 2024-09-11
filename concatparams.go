package piscine

func ConcatParams(args []string) string {
	sl := ""
	for i, v := range args {
		if i < len(args)-1 {
			sl += v + "\n"
		} else {
			sl += v
		}
	}
	return sl
}
