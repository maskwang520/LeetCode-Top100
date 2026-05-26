package main

func findTargetSumWays(nums []int, target int) int {
	result := 0
	sign := [2]int{-1, 1}
	sum := 0
	var compose func(nums []int, target int, index int)
	compose = func(nums []int, target int, index int) {
		if index < len(nums) {
			for _, s := range sign {
				sum += nums[index] * s
				compose(nums, target, index+1)
				sum -= nums[index] * s
			}
		} else {
			if sum == target {
				result++
			}
		}
	}
	compose(nums, target, 0)
	return result
}
