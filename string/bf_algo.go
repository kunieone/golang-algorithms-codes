package string

func BF(main string, sub string) int {
	length := len(sub)
	flag := -1
	i, j := 0, 0
	for i < len(main)-length+1 && j < length {
		if main[i] == sub[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
		if j == length {
			flag = i - j
		}

	}
	return flag

}
