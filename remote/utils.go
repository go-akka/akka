package remote

func reverseStringList(values []string) []string {
	lv := len(values)
	newV := make([]string, lv)
	for i := 0; i < lv; i++ {
		newV[lv-1-i] = values[i]
	}

	return newV
}
