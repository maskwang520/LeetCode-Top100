package main

func maxProfit(prices []int) int {
	list := make([]int, len(prices), len(prices))
	maxVal := 0
	for i := len(prices) - 1; i >= 0; i-- {
		maxVal = max(maxVal, prices[i])
		list[i] = maxVal
	}

	result := 0
	for i := 1; i < len(prices); i++ {
		result = max(result, list[i]-prices[i])
	}
	return result
}
