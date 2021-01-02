package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	if 2*n > size+1 {
		return false
	}

	cnt := 0
	for i, v := range flowerbed {
		if v == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == size-1 || flowerbed[i+1] == 0) {
			cnt++
			flowerbed[i] = 1
		}
	}
	return cnt >= n
}

// 简单的计数问题
// 应该也可以用总结统计, 试试看吧

// wrong
func can2(flowerbed []int, n int) bool {
	size := len(flowerbed)
	if 2*n > size+1 {
		return false
	}

	pre, cnt := -1, 0
	for i, v := range flowerbed {
		if v == 1 {
			cnt += (i - pre - 2 + 1) / 2
			pre = i
		}
	}

	cnt += (size - pre - 2 + 1) / 2
	return cnt >= n
}

// 感觉手术之后, 自己的脑子好像不怎么好使了...
// 是急了? 还是反应变慢了???

/**
整理思路:
最笨的方法: 就是暴力统计出可以种植的位置,
取巧点的: 根据规律统计
	1. 首先分析普通情况, 在两花之间种花, 此时坑位 >= 3, 才能种植
	2. 对于左|右端, 坑位 >= 2, 就能种植
	3. 统计所有两花之间的坑位, 根据情况分别计算

	问题点在于, 想要用一个方程概括各种情况, 就本题来看, 是行不通的, 或者说, 得不偿失
**/

// 先整理思路, 再coding!!!
