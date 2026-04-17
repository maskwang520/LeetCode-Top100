package main

// 方法1: 动态规划 - 2D数组
func uniquePaths(m int, n int) int {
	// 创建 dp 数组，dp[i][j] 表示到达位置 (i,j) 的路径数
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 第一行和第一列都只有 1 种路径
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 填充 dp 数组
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
