package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenOfLL := 0
	current := head
	for current != nil {
		lenOfLL++
		current = current.Next
	}
	if lenOfLL == 1 && n == 1 {
		return nil
	}
	if lenOfLL == n {
		head = head.Next
		return head
	}
	nodePositionFromStart := lenOfLL - n + 1
	var prev *ListNode
	current = head
	currNodeNo := 0
	for current != nil {
		currNodeNo++
		if currNodeNo == nodePositionFromStart {
			if current.Next != nil && prev != nil {
				prev.Next = current.Next
			} else if prev != nil {
				prev.Next = nil
			}
			break
		}
		prev = current
		current = current.Next
	}
	return head
}
