package main

func pathSum(root *TreeNode, targetSum int) int {
	prefix := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	count := 0

	dfs = func(node *TreeNode, currSum int64) {
		if node == nil {
			return
		}
		currSum += int64(node.Val)
		count += prefix[currSum-int64(targetSum)]
		prefix[currSum]++
		dfs(node.Left, currSum)
		dfs(node.Right, currSum)
		prefix[currSum]-- // 回溯，离开当前路径
	}

	dfs(root, 0)
	return count
}
