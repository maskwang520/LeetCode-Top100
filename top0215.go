package main

import "container/heap"

type minTopHeap []int

func findKthLargest(nums []int, k int) int {
	minHeap := make(minTopHeap, k)

	for i := range nums {
		if len(minHeap) > k {
			heap.Pop(&minHeap)
		}
		heap.Push(&minHeap, nums[i])
	}

	return heap.Pop(&minHeap).(int)

}

func (h *minTopHeap) Len() int {
	return len(*h)
}

func (h *minTopHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minTopHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minTopHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minTopHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
