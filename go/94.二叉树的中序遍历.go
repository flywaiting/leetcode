/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	arr := append(inorderTraversal(root.Left), root.Val)
	return append(arr, inorderTraversal(root.Right)...)
}

// 总结:
// 简单递归,
// 但是因为 golang 的特性, 传递数组的方法有问题,
// 就在直接拼接切片并返回

// 改进:
// 既然传递切片有问题 就传切片指针

func inorderTraversal2(root *TreeNode) []int {
	arr := []int{}
	inorderTraversalInner(root, &arr)
	return arr
}

func inorderTraversalInner(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	inorderTraversalInner(node.Left, arr)
	*arr = append(*arr, node.Val)
	inorderTraversalInner(node.Right, arr)
}

// @lc code=end
