package main

func searchInsert(nums []int, target int) int {
	return binarySearchFun(nums, 0, len(nums)-1, target)
}

func binarySearchFun(nums []int, left, right, target int) int {
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}
