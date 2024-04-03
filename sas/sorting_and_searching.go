package sas

import (
	"container/heap"
)

func merge(n1 []int, m int, n2 []int, n int) {
	for n != 0 {
		if m != 0 && n1[m-1] > n2[n-1] {
			n1[m+n-1] = n1[m-1]
			m--
		} else {
			n1[m+n-1] = n2[n-1]
			n--
		}
	}
}

type pq []number
type number struct {
	val, count int
}

func (nn *pq) Len() int           { return len(*nn) }
func (nn *pq) Less(i, j int) bool { return (*nn)[i].count < (*nn)[j].count }
func (nn *pq) Swap(i, j int)      { (*nn)[i], (*nn)[j] = (*nn)[j], (*nn)[i] }
func (nn *pq) Push(x any) {
	*nn = append(*nn, x.(number))
}
func (nn *pq) Pop() any {
	old := *nn
	n := len(old)
	x := old[n-1]
	*nn = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, n := range nums {
		seen[n]++
	}

	q := &pq{}
	heap.Init(q)
	for val, count := range seen {
		n := number{val: val, count: count}
		heap.Push(q, n)
		if q.Len() > k {
			_ = heap.Pop(q)
		}
	}

	out := make([]int, 0, k)
	for q.Len() > 0 {
		out = append(out, q.Pop().(number).val)
	}

	return out
}
