package leetcode

var remainArr = [...]int{0, 1, 2 * 1, 3 * 2 / 1}

func clumsy(N int) (res int) {
	idx := N % 4
	s := 1
	for N > idx {
		res += s*N*(N-1)/(N-2) + N - 3
		s = -1
		N -= 4
	}
	res += s * remainArr[idx]
	return
}

// 大概, 今天是 4/01 吧
// 就 NM 离谱, 纯数学可解
