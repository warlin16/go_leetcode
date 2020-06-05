package easy

func deleteNode(node *ListNode) {
	for node.Next != nil {
		if node.Next.Next == nil {
			break
		}
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil

}
