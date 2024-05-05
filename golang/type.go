package leetcode

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Next   *Node
	Random *Node
	Prev   *Node
	Child  *Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}
