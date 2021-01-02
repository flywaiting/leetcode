package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	cnt, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			cnt++
			right = p[1]
		}
	}
	return n - cnt
}

// 不是自己整理出来的思路...
// 首先, 需要有个解题方向, 或者说, 编程思路
// 最小去重叠, 等价于最大不重叠数量
// 无论左边的数字如何, 都是由右区间决定覆盖范围, 以此为突破口
// 然后, 就发现, 无法给出证明...

// 关于动态规划, 还有挺长的路要学习...
