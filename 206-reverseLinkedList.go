package main

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
Reverse a singly linked list.
Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/

// brute force
func reverseList(head *ListNode) *ListNode {
	stack := []*ListNode{}

	for current := head; current != nil; {
		stack = append([]*ListNode{current}, stack...)
		current = current.Next
	}
	if len(stack) == 0 || len(stack) == 1 {
		return head
	}
	newHead := stack[0]
	newHead.Next = stack[1]
	for i, current := range stack {
		if i == len(stack)-1 {
			current.Next = nil
		} else {
			current.Next = stack[i+1]
		}
	}
	return newHead
}

// optimized solution
func reverseListOptimized(head *ListNode) *ListNode {
	var newHead *ListNode
	for current := head; current != nil; {
		next := current.Next
		current.Next = newHead
		newHead = current
		current = next
	}
	return newHead
}
