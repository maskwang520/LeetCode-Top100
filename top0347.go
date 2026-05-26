package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	h := &minHeap{}
	for num, cnt := range freq {
		heap.Push(h, [2]int{num, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).([2]int)[0]
	}
	return res
}

type minHeap [][2]int // {num, freq}

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}
