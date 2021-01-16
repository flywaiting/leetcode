package leetcode

// 遍历字符串
// 记录起止区间下标
// 保存满足条件的区间
func largeGroupPositions(s string) [][]int {
	ans := [][]int{}
	cnt := 0
	for i := range s {
		if i > 0 && s[i] == s[i-1] {
			cnt++
		} else {
			if cnt > 1 {
				ans = append(ans, []int{i - cnt - 1, i - 1}) // 此时已经是下一个字符了
			}
			cnt = 0
		}
	}

	if cnt > 1 {
		n := len(s) - 1
		ans = append(ans, []int{n - cnt, n})
	}
	return ans
}

// 最后一个字符组, 需要单独处理
// (考虑不周, 漏掉了...)
