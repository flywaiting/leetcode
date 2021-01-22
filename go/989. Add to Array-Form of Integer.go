package leetcode

func addToArrayForm(A []int, K int) []int {
	ans := make([]int, 0, len(A))

	for i, ext := len(A)-1, 0; i >= 0 || K > 0 || ext > 0; i, K = i-1, K/10 {
		ext += K % 10
		if i >= 0 {
			ext += A[i]
		}
		ans = append(ans, ext%10)
		ext /= 10
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return ans
}

// 将数组转换成数字, 相加, 再转换成数组 [简单, 粗暴]
// 将 K 拆开, 融入数组
