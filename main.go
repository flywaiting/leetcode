package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello go for leetcode")

	removeDuplicateLetters2("cbacdcbc")
}

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
