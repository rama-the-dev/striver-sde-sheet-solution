package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// break list into half
	// prevOfSlow is required to create 2 separate list upto before slow one LL
	// from slow to last another LL
	prevOfSlow, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prevOfSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// create 2 separate linked lists
	prevOfSlow.Next = nil

	// do sorting on left half
	left := sortList(head)

	// do sorting on right half
	right := sortList(slow)

	// merge left and right half
	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	ansList := dummy
	var ansNode *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			ansNode = &ListNode{Val: left.Val}
			left = left.Next
		} else {
			ansNode = &ListNode{Val: right.Val}
			right = right.Next
		}
		ansList.Next = ansNode
		ansList = ansList.Next
	}
	// join left over parts of left and right to ans list
	if left != nil {
		ansList.Next = left
	}
	if right != nil {
		ansList.Next = right
	}
	return dummy.Next
}
