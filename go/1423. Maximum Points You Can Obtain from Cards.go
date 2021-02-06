package leetcode

func maxScore(cardPoints []int, k int) int {
	ans := 0
	ln := len(cardPoints)
	for _, v := range cardPoints[ln-k:] {
		ans += v
	}
	if ln == k {
		return ans
	}

	// arr := make([]int, k+1)
	pre := ans
	for i := 0; i < k; i++ {
		pre += cardPoints[i] - cardPoints[ln-k+i]
		if pre > ans {
			ans = pre
		}
	}

	return ans
}

// 思路:
// 只能从头尾抽取卡牌, 可以反向使用数组窗口
// 窗口大小为 len()-k
// 窗口移动, 根据变化计算当前值, 并取最大值

// 使用数组方便记录所有的值, 但在本题中并没有其他用途, 可以用变量保存当前值, 减少内存使用
