package leetcode

// AC
/**
正常的解题思路, 应该是用 BFS 吧,
可是, 考虑到广度需要两个队列来维护层级, 决定用 DFS 来做

然而, 在观看了别人的 BFS 之后, 这个居然内存空间更小!!!
因为用队列维护层级内容, 可以直接创建容量匹配的切片, 进而直接决定数据的排列顺序, 避免了数组翻转的操作

不管怎样, DFS 就代码方面, 还是更简洁的...
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, &res, 0)
	for i, n := 1, len(res); i < n; i += 2 {
		for j, size := 0, len(res[i]); j < size/2; j++ {
			res[i][j], res[i][size-1-j] = res[i][size-1-j], res[i][j]
		}
	}
	return res
}

func dfs(node *TreeNode, res *[][]int, idx int) {
	if node == nil {
		return
	}

	if idx >= len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[idx] = append((*res)[idx], node.Val)
	dfs(node.Left, res, idx+1)
	dfs(node.Right, res, idx+1)
}
