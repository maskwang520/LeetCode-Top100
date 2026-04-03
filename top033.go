package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] < nums[mid] {
			position := binarySearch(nums, left, mid-1, target)
			if position >= 0 {
				return position
			}
			left = mid + 1
		} else {
			position := binarySearch(nums, mid+1, right, target)
			if position >= 0 {
				return position
			}
			right = mid - 1
		}

	}
	return -1
}

func binarySearch(nums []int, left, right, target int) int {
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
	return -1
}

func main() {
	fmt.Println(search([]int{3, 1}, 3))
}
