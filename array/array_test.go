package array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		out  int
	}{
		{nums: []int{1, 1, 2}, out: 2},
		{nums: []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, out: 5},
	}

	for i, test := range tests {
		res := removeDuplicates(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}
