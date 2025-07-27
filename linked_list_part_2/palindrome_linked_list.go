package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// so the idea is to find middle point of node
	// using heir and tortoise method
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// in case fast not pointing to null
	// means LL is of odd length
	// 1 -> 2 -> 3 -> 4 -> 5 -> null
	// then move slow pointer by one move to discard odd node
	if fast != nil {
		slow = slow.Next
	}

	// reverse LL from slow to end of LL
	reversed2ndHalfOfLL := reverse_1(slow)
	firstHalfOfLL := head

	// now move both pointers simultaneously
	// upto the point 2nd LL do not exhaust
	for reversed2ndHalfOfLL != nil {
		if firstHalfOfLL.Val != reversed2ndHalfOfLL.Val {
			return false
		}
		firstHalfOfLL = firstHalfOfLL.Next
		reversed2ndHalfOfLL = reversed2ndHalfOfLL.Next
	}
	return true
}

func reverse_1(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
