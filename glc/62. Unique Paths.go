package glc

func UP62(m, n int) int {
	return uniquePaths(m, n)
}
func uniquePaths(m int, n int) int {
	dict := make([][]int, m)
	for i := range dict {
		dict[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dict[i][j] = 1
			} else if j == 0 {
				dict[i][j] = 1
			} else {
				dict[i][j] = dict[i-1][j] + dict[i][j-1]
			}
		}
	}
	return dict[m-1][n-1]
}

// 终于是个 典型、简单 的 DP 了...
// 备忘录可以压缩成 1D 数组,
// 而且, 因为旋转的等效性, 可以选择使用较小的维数作为压缩数组的长度
// 专业的叫法: 滚动数组
