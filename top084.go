package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// 栈存储索引，维护递增高度
	stack := make([]int, 0)
	maxArea := 0

	// 遍历所有柱子，包括末尾哨兵
	for i := 0; i <= n; i++ {
		// 哨兵：遍历完所有元素后，用0强制弹出栈中剩余元素
		h := 0
		if i < n {
			h = heights[i]
		}

		// 当前高度小于栈顶高度，弹出栈顶并计算面积
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			// 弹出栈顶作为矩形的高
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// 计算宽度
			// 如果栈空，说明弹出的是最小高度，宽度是 i
			// 否则宽度是 i - stack[-1] - 1
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		// 当前索引入栈
		stack = append(stack, i)
	}

	return maxArea
}
