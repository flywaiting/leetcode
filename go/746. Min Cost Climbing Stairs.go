package leetcode

// AC
// 解题思路
// 简单的DP, 用数组保存数据, 避免重复内存分配, 方面取最小值, 当然, 都是些比较细碎的优化
var need = [2]int{}

func minCostClimbingStairs(cost []int) int {
	// need := [2]int{cost[0], cost[1]}
	need[0], need[1] = cost[0], cost[1]
	n := len(cost)
	for i := 2; i < n; i++ {
		need[i%2] = min() + cost[i]
	}
	return min()
}

func min() int {
	if need[0] > need[1] {
		return need[1]
	}
	return need[0]
}

// AC
// 代码更优雅, 性能数据反而不如上面的
func minCostClimbingStairs2(cost []int) int {
	c0, c1 := 0, 0
	n := len(cost)
	for i := 0; i < n; i++ {
		c0, c1 = c1, min2(c0, c1)+cost[i]
	}
	return min()
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}
