package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}

	for _, list := range lists {
		p := dummy.Next

		cur := dummy
		for p != nil || list != nil {
			if p == nil {
				cur.Next = list
				break
			}
			if list == nil {
				cur.Next = p
				break
			}

			if p.Val > list.Val {
				cur.Next = list
				list = list.Next
			} else {
				cur.Next = p
				p = p.Next
			}
			cur = cur.Next
		}
	}

	return dummy.Next
}
