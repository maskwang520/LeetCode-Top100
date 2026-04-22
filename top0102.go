package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := make([][]int, 0)
	count := 1
	for len(queue) > 0 {
		list := make([]int, 0)
		k := 0
		for i := 0; i < count; i++ {
			head := queue[0]
			queue = queue[1:]
			if head.Left != nil {
				queue = append(queue, head.Left)
				k++
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
				k++
			}
			list = append(list, head.Val)
		}
		result = append(result, list)
		count = k
	}
	return result

}
