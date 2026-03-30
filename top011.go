package main

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxArea := 0
	for i < j {
		h1, h2 := height[i], height[j]
		maxArea = max(maxArea, min(h1, h2)*(j-i))
		if h1 > h2 {
			j--
		} else {
			i++
		}
	}
	return maxArea
}
