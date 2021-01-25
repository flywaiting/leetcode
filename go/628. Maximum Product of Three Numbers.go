package leetcode

func maximumProduct(nums []int) int {
	arr := []int{nums[0], nums[1], nums[2]}
	idx := minIdx(arr)
	for _, v := range nums[3:] {
		if v > arr[idx] {
			arr[idx] = v
			idx = minIdx(arr)
		}
	}
	return arr[0] * arr[1] * arr[2]
}

func minIdx(nums []int) int {
	idx := 0
	for i, v := range nums {
		if v < nums[idx] {
			idx = i
		}
	}
	return idx
}

// 用数组保存最大的三个值, 并记录其中最小值的下标
// 遍历原数组, 于保存数组比较
// 若大于其中的最小值, 则更新保存数组和最小值下标

// ! N/A	负负得正...
