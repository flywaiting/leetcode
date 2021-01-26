package leetcode

import "sort"

func numEquivDominoPairs(dominoes [][]int) (ans int) {
	if len(dominoes) < 2 {
		return
	}
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
	}
	sort.Slice(dominoes, func(i, j int) bool {
		if dominoes[i][0] == dominoes[j][0] {
			return dominoes[i][1] < dominoes[j][1]
		}
		return dominoes[i][0] < dominoes[j][0]
	})
	for i, v := range dominoes {
		for j := i + 1; j < len(dominoes); j++ {
			if dominoes[j][0] == v[0] && dominoes[j][1] == v[1] {
				ans++
			} else {
				break
			}
		}
	}
	return
}

// 整理数组, 以便排序
// 遍历数组, 计数满足条件的数组对

// N/A 原因
// 排序未能达到预期效果, 对 Golang 还不熟悉

// ! 总结
// 一开始的思路方向就偏了, 执着于牌是否可复用, 而忽略了其中的数学要点[相等的数量], 从而使用了最笨的方法处理[排序]

// update
// 用 map 计数相等的牌, 可直接计算处满足条件的配对数量,
// 又因为题目的约束, 可将牌转换成数组的索引

// 时间直接减半 = =

func numEquivDominoPairs2(dominoes [][]int) (ans int) {
	m := [100]int{}
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		idx := v[0]*10 + v[1]
		ans += m[idx]
		m[idx]++
	}
	return
}
