package main

func isValidBST(root *TreeNode) bool {
	list := make([]int, 0)
	travelBst(root, &list)
	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}
	return true
}

func travelBst(root *TreeNode, list *[]int) {
	if root != nil {
		travelBst(root.Left, list)
		*list = append(*list, root.Val)
		travelBst(root.Right, list)
	}
}
