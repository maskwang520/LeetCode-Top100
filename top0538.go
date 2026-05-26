package main

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right) // 先走右子树（值更大）
		sum += node.Val // 累加
		node.Val = sum  // 替换当前值
		dfs(node.Left)  // 再走左子树（值更小）
	}
	dfs(root)
	return root
}
