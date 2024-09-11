package piscine

func DescendAppendRange(max, min int) []int {
	var sl []int
	if min >= max {
		return []int{}
	}
	if max > min {
		for i := max; i > min; i-- {
			sl = append(sl, i)
		}
	}
	return sl
}
