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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	llLen := 0
	curr := head
	// store last node as this will be needed during breaking of LL
	// and pointing last node to head
	var lastNode *ListNode
	// calculate length of ll
	for curr != nil {
		llLen++
		if curr.Next == nil {
			lastNode = curr
		}
		curr = curr.Next
	}
	// find out effective ll rotation count
	effRotationCount := k
	if k == llLen {
		effRotationCount = 0
	} else if k > llLen {
		effRotationCount = k % llLen
	}
	// if no rotation required then return head as it is
	if effRotationCount == 0 {
		return head
	}
	// count breaking position from start
	breakingStartPosition := llLen - effRotationCount
	// move upto that position
	curr = head
	for i := 1; i < breakingStartPosition; i++ {
		curr = curr.Next
	}
	// break linked list at the start of breaking point
	// so current's next would become answer node or start of new linked list
	ans := curr.Next
	curr.Next = nil
	// attach last node to head
	lastNode.Next = head
	return ans
}

// func main() {
// 	ll := &ListNode{Val: 1}
// 	ans := rotateRight(ll, 0)
// 	fmt.Println(ans)
// }
