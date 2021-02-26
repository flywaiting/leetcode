package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}
	levelOrderHlp(root, 0, &arr)
	return arr
}

func levelOrderHlp(root *TreeNode, idx int, arr *[][]int) {
	if root == nil {
		return
	}

	if len(*arr) == idx {
		*arr = append(*arr, []int{})
	}
	(*arr)[idx] = append((*arr)[idx], root.Val)
	levelOrderHlp(root.Left, idx+1, arr)
	levelOrderHlp(root.Right, idx+1, arr)
}

// 解题思路:
// 1. 递归, 记录树的深度, 将值存入树深对应的数组里
