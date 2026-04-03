package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	h1, h2 := list1, list2
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}

	for h1 != nil {
		cur.Next = h1
		cur = cur.Next
		h1 = h1.Next
	}

	for h2 != nil {
		cur.Next = h2
		cur = cur.Next
		h2 = h2.Next
	}

	return newHead.Next
}
