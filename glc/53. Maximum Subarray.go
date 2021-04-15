package glc

func MaxSubArray(n []int) int {
	return maxSubArray(n)
}
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := range nums {
		if i > 0 && nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 暴力DP
// 求解各种长度的和, 选择最大值

// 自己把问题复杂化了, 另外, 还是没有抓住问题的特性, “连续”
// 首先, 是连续的, 所以只要知道 A[i] 的最大值, 对于 A[i+1], max(A[i]+ai, ai)
// 先求连续的最大值, 再选择所有连续里面的最大值
