package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	temp := head
	kthNode := &ListNode{}
	prevNode := &ListNode{}
	nextNode := &ListNode{}
	for temp != nil {
		// find kthNode
		kthNode = getKthNode(temp, k)
		if kthNode == nil {
			// if prevNode exist then attach remaining into prevNode
			if prevNode != nil {
				prevNode.Next = temp
			}
			break
		}

		// store next of kthNode into nextNode
		nextNode = kthNode.Next

		// make kthNode next to nil
		// to create complete individual LL from temp -> kthNode
		kthNode.Next = nil

		// reverse LL from temp to kthNode
		reverse(temp)

		// if this is first subgroup getting processed
		// then assign head to proper node, which will be returned as answer
		if head == temp {
			head = kthNode
		} else {
			prevNode.Next = kthNode
		}
		prevNode = temp
		temp = nextNode
	}
	return head
}

func getKthNode(head *ListNode, k int) *ListNode {
	curr := head
	nodeCount := 1
	for curr != nil && nodeCount < k {
		curr = curr.Next
		nodeCount++
	}
	return curr
}

func reverse(head *ListNode) {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
}
