package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	currAnsNode := dummy
	carryForward := 0
	currSum := 0
	currAnsNodeVal := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			currSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currSum += l2.Val
			l2 = l2.Next
		}
		currSum += carryForward
		currAnsNodeVal = currSum % 10
		carryForward = currSum / 10
		currAnsNode.Next = &ListNode{Val: currAnsNodeVal}
		currAnsNode = currAnsNode.Next
		currSum = 0
	}
	if carryForward == 1 {
		currAnsNode.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
