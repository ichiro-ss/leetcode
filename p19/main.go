package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev := head
	aftr := head
	for i := 0; i < n; i++ {
		aftr = aftr.Next
	}
	if aftr == nil {
		return prev.Next
	}
	for aftr.Next != nil {
		prev, aftr = prev.Next, aftr.Next
	}
	prev.Next = prev.Next.Next
	return head
}
