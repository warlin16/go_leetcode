package main

import "fmt"

var head *ListNode = &ListNode{
	Val:  1,
	Next: node2,
}

var node2 *ListNode = &ListNode{
	Val:  2,
	Next: node3,
}

var node3 *ListNode = &ListNode{
	Val:  3,
	Next: node4,
}

var node4 *ListNode = &ListNode{
	Val:  4,
	Next: node5,
}

var node5 *ListNode = &ListNode{
	Val:  5,
	Next: nil,
}

func main() {
	test := reverseList(head)
	for current := test; current != nil; {
		fmt.Println("This should work", current.Val)
		current = current.Next
	}
}
