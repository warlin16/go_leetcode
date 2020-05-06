package easy

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = &ListNode{}
	travel := result

	for l1 != nil || l2 != nil {
		if l1 == nil {
			travel.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			travel.Next = l1
			l1 = nil
			continue
		}
		if l1.Val > l2.Val {
			travel.Next = l2
			travel, l2 = travel.Next, l2.Next
		} else {
			travel.Next = l1
			travel, l1 = travel.Next, l1.Next
		}
	}
	return result.Next
}
