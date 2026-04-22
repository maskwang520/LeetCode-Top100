package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	deepSearchFun(root, &res)
	return res
}

func deepSearchFun(root *TreeNode, result *[]int) {
	if root != nil {
		deepSearchFun(root.Left, result)
		*result = append(*result, root.Val)
		deepSearchFun(root.Right, result)
	}
}
