package main

import "fmt"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		dp[i][0] = grid[i][0] + sum
		sum = dp[i][0]
	}

	sum = 0
	for j := 0; j < len(grid[0]); j++ {
		dp[0][j] = grid[0][j] + sum
		sum = dp[0][j]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = minVal(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
