package main

func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	currentMax := nums[0]

	for i := 1; i < len(nums); i++ {
		// 选择要么重新开始，要么继续累加
		currentMax = max(nums[i], currentMax+nums[i])
		maxSoFar = max(maxSoFar, currentMax)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
