package piscine

func Join(strs []string, sep string) string {
	var sls string
	for i := 0; i < len(strs); i++ {
		sls += strs[i]
		if strs[i] != strs[len(strs)-1] {
			sls += sep
		}
	}
	return sls
}
