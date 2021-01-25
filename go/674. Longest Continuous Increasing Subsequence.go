package leetcode

func findLengthOfLCIS(nums []int) int {
	ans, l := 0, -1
	for i := range nums {
		if i == 0 || nums[i] > nums[i-1] {
			if ans < i-l {
				ans = i - l
			}
		} else {
			l = i - 1
		}
	}
	return ans
}

// 双指针
// 遍历数组, 更新序列窗口,
