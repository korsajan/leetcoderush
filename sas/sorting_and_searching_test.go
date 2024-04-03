package sas

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var testCases = []struct {
		nums1, nums2, out []int
		m, n              int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, nums2: []int{2, 5, 6}, m: 3, n: 3, out: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, nums2: []int{}, m: 1, n: 0, out: []int{1}},
		{nums1: []int{}, nums2: []int{1}, m: 0, n: 1, out: []int{1}},
	}

	for i, test := range testCases {
		merge(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(test.nums1, test.out) {
			t.Errorf("test [%d] fail ex: %v got: %v", i+1, test.out, test.nums1)
		}
	}
}

func TestTopKFrequent(t *testing.T) {
	var testCases = []struct {
		nums, out []int
		k         int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			out:  []int{1, 2},
			k:    2,
		},
		{
			nums: []int{1},
			out:  []int{1},
			k:    1,
		},
	}

	for i, test := range testCases {
		got := topKFrequent(test.nums, test.k)
		if !reflect.DeepEqual(got, test.out) {
			t.Errorf("test [%d] fail ex: %v got: %v", i+1, test.out, got)
		}
	}
}
