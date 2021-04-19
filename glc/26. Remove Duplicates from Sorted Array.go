package glc

func RD26(n []int) int {
	return removeDuplicates2(n)
}
func removeDuplicates2(nums []int) int {
	idx := -1
	for i, v := range nums {
		if i == 0 || v != nums[idx] {
			idx++
			nums[idx] = v
		}
	}
	return idx + 1
}
