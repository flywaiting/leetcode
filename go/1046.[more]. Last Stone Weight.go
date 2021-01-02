package leetcode

import (
	"sort"
)

// AC
/**
消除元素, 排序, 简单粗暴的做法, 还以为有什么其他更好的方式可以进行动态排序

利用堆栈好像可以会更简便, 是因为堆栈自身在添加元素的时候, 进行了排序 ???
golang: heap ???

**/
func lastStoneWeight(stones []int) int {
	n := len(stones)
	for n > 1 {
		sort.Ints(stones)
		x, y := stones[n-2], stones[n-1]
		if x == y {
			stones = stones[:n-2]
			n -= 2
		} else {
			stones[n-2] = y - x
			stones = stones[:n-1]
			n--
		}
	}

	if n == 0 {
		return 0
	}
	return stones[0]
}
