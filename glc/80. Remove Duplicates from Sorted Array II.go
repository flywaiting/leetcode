package glc

func RD80(n []int) int {
	return removeDuplicates(n)
}
func removeDuplicates(nums []int) int {
	ln, cnt := 0, 0
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			nums[ln] = nums[i]
			ln++
		}
	}
	return ln
}
