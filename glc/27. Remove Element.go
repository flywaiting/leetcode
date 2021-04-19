package glc

func RE27(n []int, v int) int {
	return removeElement(n, v)
}
func removeElement(nums []int, val int) int {
	idx := -1
	for _, v := range nums {
		if v != val {
			idx++
			nums[idx] = v
		}
	}
	return idx + 1
}
