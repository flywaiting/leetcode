package leetcode

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for head != fast {
				head = head.Next
				fast = fast.Next
			}
			return head
		}
	}
	return nil

	// for head != nil {
	// 	if head.Val == 1e6 {
	// 		return head
	// 	}
	// 	head.Val = 1e6
	// 	head = head.Next
	// }
	// return nil
}
