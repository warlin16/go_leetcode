package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	copiedHead := &ListNode{Next: head}

	for n > 0 {
		head = head.Next
		n--
	}
	previous := copiedHead
	for head != nil {
		head, previous = head.Next, previous.Next
	}

	previous.Next = previous.Next.Next
	return copiedHead.Next
}
