package leetcode

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	cnt, idx := 0, 0
	for i, v := range nums {
		if i == 0 || v != nums[i-1] {
			cnt = 0
		} else {
			cnt++
		}

		if cnt < 2 {
			nums[idx] = v
			idx++
		}
	}
	return idx
}

// 遍历数组
// 计数相等数值
// 最多只保留2个相等数值
// 不允许新开数组, 就原地修改, 需要额外的下标记录整理之后的数组

// 奇淫巧计
// 利用排序数组, 最多2个相同值, 必定有当前值大于重写数组的的倒数第二个数
// 如果当前值不大于整理后的倒数第二个数, 说明相同的数重复出现了2次以上
// 此时不应该继续添加相同数值了

func learn(nums []int) int {
	idx := 0
	for _, v := range nums {
		if idx < 2 || v > nums[idx-2] {
			nums[idx] = v
			idx++
		}
	}
	return idx
}
