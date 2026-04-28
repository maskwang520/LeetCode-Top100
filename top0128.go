package main

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, n := range nums {
		set[n] = true
	}

	best := 0
	for n := range set {
		// 只从序列起点开始统计(n-1 不存在时,n 才是起点)
		if set[n-1] {
			continue
		}
		cur := n
		length := 1
		for set[cur+1] {
			cur++
			length++
		}
		if length > best {
			best = length
		}
	}
	return best
}
