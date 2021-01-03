package leetcode

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{}
	sufHead := &ListNode{}
	pre, suf := preHead, sufHead
	for head != nil {
		if head.Val < x {
			pre.Next = head
			pre = pre.Next
		} else {
			suf.Next = head
			suf = suf.Next
		}

		head = head.Next
	}
	pre.Next = sufHead.Next
	suf.Next = nil
	return preHead.Next
}

// 遍历列表
// 比较节点数值, 将其分别保存
// 整合两个列表, 结果即为所求
