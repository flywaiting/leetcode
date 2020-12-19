package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}
	// 对角线
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 中线
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		for row := 0; row < n; row++ {
			matrix[row][i], matrix[row][j] = matrix[row][j], matrix[row][i]
		}
	}
}

/**
解题思路:
利用两次对折实现矩阵的旋转, 这个算是矩阵自身的特性吧

其他思路:
从一般向普适推广, 通过观察可以发现, 因为四条边的关系, 一个位置的旋转, 有4个数与其关联, 一次做这4个数的移位, 就能实现旋转了
*/
