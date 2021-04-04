package glc

import "sort"

func NumRabbits(a []int) int {
	return numRabbits(a)
}

func numRabbits(answers []int) (res int) {
	sort.Ints(answers)

	limit, cnt := 1, 0
	for i, v := range answers {
		// same
		if i > 0 && v == answers[i-1] && limit > cnt {
			cnt++
		} else {
			limit = v + 1
			cnt = 1
			res += limit
		}
	}
	return
}

// 简单的集合问题,
// 注意相同数量和最大限制数量, ep: 5个3 需要分两组
