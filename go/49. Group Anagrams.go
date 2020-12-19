package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	dict := make(map[string][]string)
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		ss := string(b)
		dict[ss] = append(dict[ss], s)
	}
	res := make([][]string, 0, len(dict))
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

/**
1. 寻找特征码
2. 用特征码标记、归类

鬼才级想法：利用素数不可约分的特点，将字母映射到素数，然后涌起乘积作为特征码
我的想法：用字符串本身作为特征码，因为乱序，需要对字符串进行排序；
	【不可行的想法】：利用bite位记录字母，作为特征码，可是字符串的字符有重复的可能，故不成立
	以此推进，利用数组作为特征码，用数字计数，可以无数字符串顺序，又能记录重复信息
*/
