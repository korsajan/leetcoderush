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

func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		out  []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, out: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, out: []int{3, 99, -1, -100}},
	}

	for i, test := range tests {
		rotate(test.nums, test.k)
		if !reflect.DeepEqual(test.out, test.nums) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, test.nums)
		}
	}
}

func TestContainsDuplicate(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  bool
	}{
		{nums: []int{1, 2, 3, 1}, out: true},
		{nums: []int{1, 2, 3, 4}, out: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, out: true},
	}

	for i, test := range tableTest {
		res := containsDuplicate(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestSingleNumber(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  int
	}{
		{nums: []int{2, 2, 1}, out: 1},
		{nums: []int{4, 1, 2, 1, 2}, out: 4},
		{nums: []int{1}, out: 1},
	}

	for i, test := range tableTest {
		res := singleNumber(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}
