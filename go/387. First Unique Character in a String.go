package leetcode

/**
点滴:
关注需求, 对于对于非目标数据, 可以二次利用
*/

// ac
/**
解题思路:
简单的记录, 定位
不想进行两次字符串遍历, 就用额外的切片保存字符出现顺序和下标信息
*/
func firstUniqChar(s string) int {
	cntArr := make([]int, 26, 26)
	seq := []int{}
	for i, ch := range s {
		if cntArr[ch-'a'] < 1 {
			seq = append(seq, i*100+int(ch-'a'))
		}
		cntArr[ch-'a']++
	}
	for _, i := range seq {
		if cntArr[i%100] == 1 {
			return i / 100
		}
	}
	return -1
}

/**
学习:
用一个数组记录字符出现的下标, 对于重复出现的字符, 将其置为特殊值, 因为需求只需要关注唯一出现的字符
然后遍历数组, 找出最小下标值, 即为所求
*/
func firstUniqChar2(s string) int {
	seq := [26]int{}
	n := len(s)
	for i, ch := range s {
		if seq[ch-'a'] == 0 {
			seq[ch-'a'] = i + 1
		} else {
			seq[ch-'a'] = -1
		}
	}
	ans := n + 1
	for _, i := range seq {
		if seq[i] > 0 && seq[i] < ans {
			ans = seq[i]
		}
	}
	if ans < n+1 {
		return ans - 1
	}
	return -1
}
