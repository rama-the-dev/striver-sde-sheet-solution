package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// algo
	// take 2 pointers, slow and fast
	// move fast pointer n position ahead to create gap of 2
	// now move both slow and fast pointer upto the point fast.Next == nil
	// now we have to remove slow.Next
	noOfStepsMoved := 0
	slow, fast := dummy, dummy
	for noOfStepsMoved != n && fast != nil {
		fast = fast.Next
		noOfStepsMoved++
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
