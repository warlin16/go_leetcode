package medium

import "math"

// ListNode declaration
type ListNode struct {
	Val  int
	Next *ListNode
}

var v = 0
var c *int = &v

// AddTwoNumbers (2 -> 4 -> 3) + (5 -> 6 -> 4) = (7 -> 0 -> 8)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val := l1.Val + l2.Val + *c
	*c = int(math.Floor(float64(val / 10)))
	ret := ListNode{val % 10, nil}

	if l1.Next != nil || l2.Next != nil || *c != 0 {
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{0, nil}
		}
		ret.Next = AddTwoNumbers(l1.Next, l2.Next)
	}
	return &ret
}
