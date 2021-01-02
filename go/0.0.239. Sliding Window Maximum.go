package leetcode

// 超时
// 更新窗口最大值的
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	n := len(nums)
	arr := make([]int, n-k+1)
	for i, v := range nums[k-1:] {
		if i > 0 && v >= arr[i-1] {
			arr[i] = v
		} else {
			arr[i] = max(nums[i : i+k])
		}
	}
	return arr
}

func max(nums []int) int {
	ans := nums[0]
	for _, v := range nums {
		if ans < v {
			ans = v
		}
	}
	return ans
}

// 通过维护一个可以立刻获取最大值的队列, 就可以避免重复计算了
// 队列维护:
// 1. 将小于待处理的值清空;
// 2. 将超窗的最大值清理掉;
// 3. 将待处理值加入队列;
// 4. 获取最大值
