package glc

func trap2(height []int) int {
	pre, res := 0, 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			res += add(pre - height[i])
			pre = max(pre, height[i])
			i++
		} else {
			res += add(pre - height[j])
			pre = max(pre, height[j])
			j--
		}
	}
	return res
}
