package leetcode

import "sort"

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int)
	for _, v := range nums {
		if _, has := dict[v]; has {
			return true
		}
		dict[v] = 1
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

/**
1. 使用 if-else 比 if-return 更快，猜测 golang 有进行优化
2. 在内存方面，使用 int 标识，比使用 bool 更省内存， => 网络传输上，也是如此？？？
*/
