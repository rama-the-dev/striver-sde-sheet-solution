package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input 1 -> 2 -> 3 -> 4 -> 5
// algo
// focus is to write code which is iterable with same logic upto the end of LL
// 1. take dummy prev node
// on first iteration
// 2. store curr.Next in next
// 3. break link of curr by pointing curr to prev (which would be nil in first iteration)
// 4. now move prev and head to it's respective next
// i.e. prev = head and head = next
// now repeat 1->4 upto the point head is not nil

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
