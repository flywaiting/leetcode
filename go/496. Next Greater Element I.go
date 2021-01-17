package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m[v] = -1
	}

	n := len(nums2)
	for i, max := n-2, nums2[n-1]; i >= 0; i-- {
		if _, has := m[nums2[i]]; has && nums2[i] < max {
			for j := i + 1; j < n; j++ {
				if nums2[j] > nums2[i] {
					m[nums2[i]] = nums2[j]
					break
				}
			}
		}
		if nums2[i] > max {
			max = nums2[i]
		}
	}

	for i := range nums1 {
		nums1[i] = m[nums1[i]]
	}
	return nums1
}

// 整理数组, 求解元素的目标值
// 从右往左, 记录出现的最大值, 方便判定是否存在目标值
// 以子数组过滤需要征高丽的元素

// 进阶
// 同样是先整理数组,
// 使用单调栈求解目标值
// 将目标值用map保存, 之后遍历子数组, 将对应数值取出

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums2))
	stack := []int{}
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			m[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	for _, v := range stack {
		m[v] = -1
	}

	for i, v := range nums1 {
		nums1[i] = m[v]
	}
	return nums1
}
