package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// algo 1
	// store copy of each node into a map
	nodeCopyMap := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		copyNode := &Node{Val: curr.Val}
		nodeCopyMap[curr] = copyNode
		curr = curr.Next
	}
	// now assign next and random pointers to copy nodes
	curr = head
	for curr != nil {
		copyNode := nodeCopyMap[curr]
		copyNode.Next = nodeCopyMap[curr.Next]
		copyNode.Random = nodeCopyMap[curr.Random]
		curr = curr.Next
	}
	return nodeCopyMap[head]
}

func copyRandomList_2(head *Node) *Node {
	// algo 2
	// first iteration : insert copy nodes, next to original node
	// second iteration : insert next and random pointers to copy nodes
	// thirst iteration : separate out the original and copy linked lists
	curr := head
	for curr != nil {
		copyNode := &Node{Val: curr.Val, Next: curr.Next}
		curr.Next = copyNode
		curr = copyNode.Next
	}
	curr = head
	for curr != nil {
		copyNode := curr.Next
		if curr.Random != nil {
			copyNode.Random = curr.Random.Next
		}
		curr = copyNode.Next
	}
	ans := &Node{}
	dummy := ans
	curr = head
	for curr != nil {
		// creating new list
		ans.Next = curr.Next
		ans = ans.Next

		// disconnecting and going back to original list
		curr.Next = curr.Next.Next
		curr = curr.Next
	}
	return dummy.Next
}
