package string

import "math"

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

func myAtoi(s string) int {
	i := 0
	// trim whitespace
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}

	if i == len(s) {
		return 0
	}

	var neg = 1
	switch s[i] {
	case '-':
		neg = -1
		i++
	case '+':
		i++
	}

	n := 0
	b := []byte(s[i:])
	for _, ch := range b {
		ch -= '0'
		if ch > 9 {
			break
		}
		n = n*10 + int(ch)
		if n*neg > math.MaxInt32 {
			return math.MaxInt32
		}
		if n*neg < math.MinInt32 {
			return math.MinInt32
		}
	}

	return n * neg
}
