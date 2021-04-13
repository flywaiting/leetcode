package glc

import "math"

var ans, pre int

func minDiffInBST(root *TreeNode) int {
	ans, pre = math.MaxInt64, -1
	minDiffInBSTDFS(root)
	return ans
}

func minDiffInBSTDFS(root *TreeNode) {
	if root == nil {
		return
	}

	minDiffInBSTDFS(root.Left)
	if pre >= 0 {
		ans = min(ans, root.Val-pre)
	}
	pre = root.Val
	minDiffInBSTDFS(root.Right)
}

// 二叉搜索树, 中序遍历
// 二叉树相关的东西, 还是了解太少了...
