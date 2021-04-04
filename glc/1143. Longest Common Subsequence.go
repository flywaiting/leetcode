package glc

import "fmt"

func LongestCommonSubsequence(s1, s2 string) int {
	return longestCommonSubsequence2(s1, s2)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	row, col := len(text1), len(text2)
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}

	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[row][col]
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	if len(text2) > len(text1) {
		text1, text2 = text2, text1
	}

	n1 := len(text1)
	n2 := len(text2)

	dp := make([]int, n1+1)
	for i := 1; i <= n1; i++ {
		preDP := 0
		for j := 1; j <= n2; j++ {
			var val int
			if text1[i-1] == text2[j-1] {
				val = preDP + 1
				fmt.Println("T", val, preDP)
			} else {
				val = max(dp[j-1], dp[j])
				fmt.Println("F", val, preDP)
			}
			preDP = dp[j]
			dp[j] = val
			fmt.Println(val)
		}
	}

	return dp[n2]
}

// DP 还是不熟, 需要系统、深入的学习!!
// 典型的 DP 问题, 还不熟悉, 非遍历方向找不到思路...
// 压缩数据, 只保留当前有效数据, 可以节约内存, 解题上称为 滚动压缩
// 压缩需要格外管理之前的最大值, 以更新、比较
