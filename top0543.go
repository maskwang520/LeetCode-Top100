package main

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		// 经过当前节点的路径长度 = 左深度 + 右深度
		ans = max(ans, left+right)
		// 返回当前节点的深度
		return max(left, right) + 1
	}

	depth(root)
	return ans
}
