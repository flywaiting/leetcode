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

func add(n int) int {
	if n < 0 {
		return 0
	} else {
		return n
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 双指针, 一次遍历
