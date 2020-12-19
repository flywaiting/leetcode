package leetcode

func findTheDifference(s string, t string) byte {
	arr := [26]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		arr[s[i]-'a']--
		arr[t[i]-'a']++
	}
	arr[t[n]-'a']++
	var b byte
	for i := 0; i < 26; i++ {
		if arr[i] > 0 {
			b = byte(i) + 'a'
			break
		}
	}
	return b
}

func findTheDifference2(s string, t string) byte {
	var b byte
	n := len(s)
	for i := 0; i < n; i++ {
		b ^= s[i] ^ t[i]
	}
	return b ^ t[n]
}

/**
解题思路
	简单的计数题目

其他思路:
	多出来的字母, 单独存在, 奇数次, 异或求值
*/
