package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	nowNode := newHead
	flag := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + flag
		if val > 9 {
			val = val % 10
			flag = 1
		} else {
			flag = 0
		}
		node := ListNode{Val: val}
		nowNode.Next = &node
		nowNode = &node
		l2 = l2.Next
		l1 = l1.Next
	}

	for l1 != nil {
		val := l1.Val + flag
		if val > 9 {
			val = val % 10
			flag = 1
		} else {
			flag = 0
		}
		node := ListNode{Val: val}
		nowNode.Next = &node
		nowNode = &node
		l1 = l1.Next
	}

	for l2 != nil {
		val := l2.Val + flag
		if val > 9 {
			val = val % 10
			flag = 1
		} else {
			flag = 0
		}
		node := ListNode{Val: val}
		nowNode.Next = &node
		nowNode = &node
		l2 = l2.Next
	}

	if flag == 1 {
		node := ListNode{Val: 1}
		nowNode.Next = &node
	}
	nowNode = nil
	return newHead.Next
}
