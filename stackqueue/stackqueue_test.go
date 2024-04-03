package stackqueue

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	var testCases = []struct {
		in, out string
	}{
		{in: "3[a]2[bc]", out: "aaabcbc"},
		{in: "3[a2[c]]", out: "accaccacc"},
		{in: "2[abc]3[cd]ef", out: "abcabccdcdcdef"},
	}

	for i, test := range testCases {
		got := decodeString(test.in)
		if got != test.out {
			t.Errorf("test %d fail ex: '%s' got: '%s'", i+1, test.out, got)
		}
	}
}

func TestIsValid(t *testing.T) {
	var testCases = []struct {
		in  string
		out bool
	}{
		{in: "()", out: true},
		{in: "()[]{}", out: true},
		{in: "(]", out: false},
		{in: "((()))", out: true},
	}

	for i, test := range testCases[3:] {
		got := isValid(test.in)
		if got != test.out {
			t.Errorf("test %d fail ex: '%t' got: '%t'", i+1, test.out, got)
		}
	}
}
