package leetcode

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if nums[0] >= 0 {
		return nums[n-1] * nums[n-2] * nums[n-3]
	}

	h := nums[0] * nums[1] * nums[n-1]
	t := nums[n-1] * nums[n-2] * nums[n-3]
	if h > t {
		return h
	}
	return t
}

// N/A
// 获取数组的前3个值, 然后, 就错了
// 问题点: 负负得正

// 暴力
// 排序, 取前后的最大值

// 提升
// 3个最大值的基础上, 收集2个最小值		=>	线性扫描
