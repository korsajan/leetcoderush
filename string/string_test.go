package string

import "testing"

func TestStrStr(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		out      int
	}{
		{haystack: "sadbutsad", needle: "sad", out: 0},
		{haystack: "leetcode", needle: "leeto", out: -1},
	}

	for i, test := range testCases {
		res := strStr(test.haystack, test.needle)
		if res != test.out {
			t.Errorf("test [%d] failed ex: %d want: %d", i+1, test.out, res)
			return
		}
	}
}

func TestFirstUniqChar(t *testing.T) {
	testCases := []struct {
		input string
		out   int
	}{
		{input: "leetcode", out: 0},
		{input: "loveleetcode", out: 2},
		{input: "aabb", out: -1},
	}

	for i, test := range testCases {
		res := firstUniqChar(test.input)
		if res != test.out {
			t.Errorf("test [%d] failed ex: %d want: %d", i+1, test.out, res)
			return
		}
	}
}
