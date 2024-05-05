package leetcode

func flatten(root *Node) *Node {
	deepFlatten(root)
	return root
}

func deepFlatten(root *Node) *Node {
	for root != nil {
		next := root.Next
		if root.Child != nil {
			last := deepFlatten(root.Child)
			if next == nil {
				next = last
			} else {
				last.Next, next.Prev = next, last
			}
			root.Next, root.Child.Prev = root.Child, root
			root.Child = nil
		}

		if next == nil {
			break
		}
		root = next
	}
	return root
}
