package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make([]int, 128)
	ans, head := 0, 0
	for i, c := range s {
		if m[c] > head {
			head = m[c]
		}
		if ans < i-head+1 {
			ans = i - head + 1
		}
		m[c] = i + 1
	}
	return ans
}

// 记录字符下标
// 重复出现的字符判定是否在子串内, 若在子串内, 则更新子串的起始点
// 计算当前子串长度, 取最大子串长度
// 更新字符下标值

// 不能记录字符下标, 而是字符下标next, 用以区分以当前下标为起始点的子串
