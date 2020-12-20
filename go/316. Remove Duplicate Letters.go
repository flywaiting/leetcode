package leetcode

import (
	"fmt"
	"sort"
)

// NA
// 难点1: 需要保留原来的字符串相对位置
// 难点2: 最小结果
func removeDuplicateLetters(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	arr := make([]int, 26)
	var c byte
	for i := n - 1; i >= 0; i-- {
		if arr[s[i]-'a'] == 0 || s[i] < c {
			c = s[i]
			arr[s[i]-'a'] = i + 1
		}
	}

	dict := make(map[int]int)
	for i := 0; i < 26; i++ {
		if arr[i] > 0 {
			dict[arr[i]] = i
		}
	}

	sort.Ints(arr)
	b := []byte{}
	for i := 0; i < 26; i++ {
		if arr[i] > 0 {
			b = append(b, byte(dict[arr[i]])+'a')
		}
	}
	return string(b)
}

// AC
// 学习要点
// 1. 排列, 尽可能将小的排前面
// 2. 动态追踪已有的字母, 增减时进行变更
// [错误] 3. 特殊处理不重复字母
// 4. 注意剩余字符
func removeDuplicateLetters2(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	arr := [26]int{}
	for i := 0; i < n; i++ {
		arr[s[i]-'a']++
	}

	res := make([]byte, 0, 26)
	rcdList := [26]int{}
	for i := 0; i < n; i++ {
		last := len(res) - 1
		for last >= 0 && rcdList[s[i]-'a'] == 0 && arr[res[last]-'a'] > 0 && res[last] > s[i] {
			rcdList[res[last]-'a'] = 0
			res = res[:last]
			last--
		}

		fmt.Println(last, string(s[i]), string(res))

		arr[s[i]-'a']--
		if last < 0 || rcdList[s[i]-'a'] == 0 {
			res = append(res, s[i])
			rcdList[s[i]-'a'] = 1
		}
		fmt.Println(string(res))
	}
	return string(res)
}

// 优化
// 1. 只有还未添加的字符串, 才有处理的必要
// 2. 弹出条件, 有内容可弹出, 当前字符大于待处理字符, 当前字符还有重复字符
func removeDuplicateLetters3(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	arr := [26]int{}
	for i := 0; i < n; i++ {
		arr[s[i]-'a']++
	}

	res := make([]byte, 0, 26)
	rcdList := [26]int{}
	for i := 0; i < n; i++ {
		if rcdList[s[i]-'a'] == 0 {
			last := len(res) - 1
			for last >= 0 && arr[res[last]-'a'] > 0 && res[last] > s[i] {
				rcdList[res[last]-'a'] = 0
				res = res[:last]
				last--
			}
			res = append(res, s[i])
			rcdList[s[i]-'a'] = 1
		}
		arr[s[i]-'a']--
	}
	return string(res)
}
