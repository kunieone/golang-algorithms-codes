package string

func getNext(b string, next []int, lenb int) {
	if lenb == 1 {
		next[0] = -1
		return
	}
	(next)[0] = -1
	(next)[1] = 0
	j := 2
	k := 0

	for j < lenb {
		if k != -1 || b[j-1] == b[k] {
			(next)[j] = k + 1
			j++
			k++
		} else {
			k = (next)[k]
		}
	}

}

func KMP(a string, b string) int {
	i := 0
	j := 0
	lena := len(a)
	lenb := len(b)
	next := make([]int, lenb)
	getNext(b, next, lenb)
	for i < lena && j < lenb {
		if j == -1 || a[i] == b[j] {
			i++
			j++

		} else {
			j = next[j]
		}
	}
	if j >= lenb {
		return i - j
	}
	next = nil
	return -1
}
