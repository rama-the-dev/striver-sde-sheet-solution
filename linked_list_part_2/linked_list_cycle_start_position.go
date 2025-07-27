package main

import "fmt"

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

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func main() {
	// node1 := &ListNode{Val: 3}
	// node4 := &ListNode{Val: 4, Next: node1}
	// node3 := &ListNode{Val: 0, Next: node4}
	// node2 := &ListNode{Val: 2, Next: node3}
	// node1.Next = node2

	node1 := &ListNode{Val: 1}

	ans := detectCycle(node1)
	fmt.Println(ans)
}
