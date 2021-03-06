package glc

func Trap(h []int) int {
	return trap(h)
}

func trap(height []int) (res int) {
	pre := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			res += add(pre - height[i])
			pre = max(height[i], pre)
			i++
		} else {
			res += add(pre - height[j])
			pre = max(pre, height[j])
			j--
		}
	}
	return
}

// 双指针, 一次遍历
