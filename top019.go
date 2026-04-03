package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tempHead := &ListNode{Next: nil}
	tempHead.Next = head
	h := tempHead.Next
	count := 0
	for h != nil {
		count++
		h = h.Next
	}

	c := 0
	h, pre := tempHead, tempHead
	for c <= count-n {
		pre = h
		h = h.Next
		c++
	}
	pre.Next = pre.Next.Next
	return tempHead.Next
}
