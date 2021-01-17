package leetcode

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}

	a, c := coordinates[0][1]-coordinates[1][1], coordinates[0][0]-coordinates[1][0]
	b := c*coordinates[0][1] - a*coordinates[0][0]
	for _, v := range coordinates[2:] {
		if c*v[1] != a*v[0]+b {
			return false
		}
	}
	return true
}

// 简单的线性方程
// 需要注意浮点数问题

// 死板的线性方法
// 向量更方便, 构造法向量, 然后为所欲为...

func checkStraightLine2(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}

	px, py := -(coordinates[0][1] - coordinates[1][1]), coordinates[0][0]-coordinates[1][0]
	for _, v := range coordinates[2:] {
		if px*(v[0]-coordinates[0][0])+py*(v[1]-coordinates[0][1]) != 0 {
			return false
		}
	}
	return true
}
