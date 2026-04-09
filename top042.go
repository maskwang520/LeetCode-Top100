package main

import "fmt"

func trap(height []int) int {
	result := 0
	leftSlice := make([]int, len(height))
	rightSlice := make([]int, len(height))
	maxLeft := 0
	for i := 1; i < len(height); i++ {
		maxLeft = max(maxLeft, height[i-1])
		leftSlice[i] = maxLeft
	}

	maxRight := 0
	for i := len(height) - 2; i >= 0; i-- {
		maxRight = max(maxRight, height[i+1])
		rightSlice[i] = maxRight
	}

	for i := 1; i < len(height); i++ {
		area := min(leftSlice[i], rightSlice[i]) - height[i]
		if area > 0 {
			result += area
		}
	}

	return result

}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
