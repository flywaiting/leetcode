package glc

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	n := len(s)
	dict := make([][]bool, n)
	for i := range dict {
		dict[i] = make([]bool, n)
	}

	start, lng := 0, 0
	for l := range s { // 设定回文长度
		for i := 0; i+l < n; i++ { // 索引
			if l == 0 {
				dict[l][i] = true
			} else if l == 1 {
				dict[l][i] = s[i] == s[i+1]
			} else {
				dict[l][i] = s[i] == s[i+l] && dict[l-2][i+1]
			}

			if dict[l][i] && l >= lng {
				start, lng = i, l+1
			}
		}
	}
	return s[start : start+lng]
}

// DP
// 暴力求解所有可能长度的回文,
// 利用表记录, 避免二次计算

// 中心扩展
// 遍历字符串, 从索引位置扩展, 确定以此为中心的回文
// 对连续重复字符特殊处理
