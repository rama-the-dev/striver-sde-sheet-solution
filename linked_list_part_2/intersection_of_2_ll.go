package main

import "math"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1Len := 0
	l2Len := 0

	currA := headA
	for currA != nil {
		l1Len++
		currA = currA.Next
	}

	currB := headB
	for currB != nil {
		l2Len++
		currB = currB.Next
	}

	currA, currB = headA, headB
	lenDiff := int(math.Abs(float64(l1Len - l2Len)))
	if l1Len > l2Len {
		for i := 1; i <= lenDiff; i++ {
			currA = currA.Next
		}
	} else if l2Len > l1Len {
		for i := 1; i <= lenDiff; i++ {
			currB = currB.Next
		}
	}
	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
	}
	return nil
}

func getIntersectionNode_2(headA, headB *ListNode) *ListNode {
	l1Len := 0
	l2Len := 0

	currA := headA
	for currA != nil {
		l1Len++
		currA = currA.Next
	}

	currB := headB
	for currB != nil {
		l2Len++
		currB = currB.Next
	}

	currA, currB = headA, headB
	for l1Len > l2Len {
		l1Len--
		currA = currA.Next
	}
	for l2Len > l1Len {
		l2Len--
		currB = currB.Next
	}
	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
	}
	return nil
}
