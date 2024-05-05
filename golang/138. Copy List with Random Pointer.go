package leetcode

func copyRandomList(head *Node) *Node {
	p := &Node{}
	tail := p
	cur := head

	i := 0
	list := []*Node{}
	for cur != nil {
		node := &Node{Val: cur.Val}
		tail.Next = node
		tail = tail.Next

		cur.Val = i
		i++
		cur = cur.Next

		list = append(list, node)
	}

	cur = head
	tail = p.Next
	for cur != nil {
		if cur.Random != nil {
			tail.Random = list[cur.Random.Val]
		}
		cur = cur.Next
		tail = tail.Next
	}

	return p.Next
}

func better(head *Node) *Node {
	if head == nil {
		return nil
	}

	p := head
	for p != nil {
		node := &Node{Val: p.Val}
		node.Next = p.Next
		p.Next = node
		p = node.Next
	}

	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	res := &Node{}
	cur := res
	p = head
	for p != nil {
		cur.Next = p.Next
		p.Next = p.Next.Next
		cur = cur.Next
		p = p.Next
	}
	return res.Next
}

// 解题思路
// 第一次
// 先处理复制 然后记录对应位置 一一映射过去
// 然后 就超时了。。。

// 第二次
// 修改数据 当索引用 这样只需要处理一个数组
// 其实 用哈希表更简单 也不用修改数据

// 提升
// 根据链表的特性 先将复制数据接在原始链表中 并用原始链表充当索引 找到随机指针对应的节点 根据节点找到复制节点 然后填充复制数据的随机指针
// 节省为了映射额外的空间开销
