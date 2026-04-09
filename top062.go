package main

func uniquePaths(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// currentSum 表示当前子数组的和
	// maxSum 表示全局最大子数组和
	currentSum := nums[0]
	maxSum := nums[0]

	// 从第二个元素开始遍历
	for i := 1; i < len(nums); i++ {
		// 如果当前和为负数，则从当前元素重新开始
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			// 否则将当前元素加到当前和中
			currentSum += nums[i]
		}

		// 更新最大和
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
