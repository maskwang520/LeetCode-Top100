package main

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixCount := map[int]int{0: 1} // 前缀和为0出现1次（空前缀）

	for _, num := range nums {
		sum += num
		if c, ok := prefixCount[sum-k]; ok {
			count += c
		}
		prefixCount[sum]++
	}

	return count
}
