package main

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 负贡献直接丢弃
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)

		// 以当前节点为拐点的完整路径
		cur := node.Val + left + right
		if cur > maxSum {
			maxSum = cur
		}

		// 向上返回时只能选一边
		return node.Val + max(left, right)
	}

	dfs(root)
	return maxSum
}
