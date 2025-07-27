package main

func hasCycle_do_not_work(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle_works(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// func main() {
// 	head := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 				},
// 			},
// 		}}
// 	ans := hasCycle_works(head)
// 	fmt.Println(ans)
// }
