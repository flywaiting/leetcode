package leetcode

func insertionSortList(head *ListNode) *ListNode {
	res := &ListNode{}
	for head != nil {
		for cur := res; true; cur = cur.Next {
			if cur.Next == nil || cur.Next.Val > head.Val {
				head, cur.Next, head.Next = head.Next, head, cur.Next
				break
			}
		}
	}

	return res.Next
}
