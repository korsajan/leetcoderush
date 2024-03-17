package string

func strStr(haystack string, needle string) int {
	ln := len(needle)

	for i := 0; i < len(haystack); i++ {
		if ln > len(haystack) {
			ln = len(haystack)
		}
		subStr := haystack[i:ln]
		if subStr == needle {
			return i
		}
		ln++
	}
	return -1
}

func firstUniqChar(ss string) int {
	counting := make([]int, 26)
	for _, s := range ss {
		counting[int(s-'a')]++
	}
	for i, s := range ss {
		if counting[int(s-'a')] == 1 {
			return i
		}
	}
	return -1
}
