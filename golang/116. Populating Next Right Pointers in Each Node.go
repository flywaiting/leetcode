package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	list := []*Node{root}
	for len(list) > 0 {
		next := []*Node{}
		for i, v := range list {
			if i < len(list)-1 {
				v.Next = list[i+1]
			}
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		list = next
	}
	return root
}
