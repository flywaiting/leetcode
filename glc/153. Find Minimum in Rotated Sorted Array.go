package glc

func FM153(n []int) int {
	return findMin((n))
}
func findMin(nums []int) int {
	for mid, i, j := 0, 0, len(nums)-1; ; {
		mid = (i + j) / 2
		if nums[i] < nums[mid] { // 左边顺序
			if nums[i] < nums[j] {
				j = mid
			} else {
				i = mid
			}
		} else {
			j = mid
		}

		if j-i <= 1 {
			return min(nums[i], nums[j])
		}
	}
}

// binary search

// upgrade
// 用差值区间求中值, 可以避免各种情况的判定
