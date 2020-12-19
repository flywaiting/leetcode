package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	strArr := strings.Fields(s)
	if len(pattern) != len(strArr) {
		return false
	}

	pm := map[rune]string{}
	sm := map[string]int{}
	for idx, c := range pattern {
		if s, has := pm[c]; has && s != strArr[idx] {
			return false
		}
		pm[c] = strArr[idx]
		sm[strArr[idx]] = 0
	}
	return len(pm) == len(sm)
}

/**
解题思路：
两个集合【双射】，相互映射
*/
