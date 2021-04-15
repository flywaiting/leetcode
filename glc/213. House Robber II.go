package glc

// func rob(nums []int) int {
// 	ans := [2]int{}
// 	n := len(nums)

// 	ans[0] = nums[0]
// 	if n>=2 {
// 		ans[1] = nums[1]
// 	}

// 	f := make([]int, n)
// 	for i:=2;i<n;i++ {

// 	}
// }

func R213(n []int) int {
	return rob(n)
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(robMinu(nums[:n-1]), robMinu(nums[1:]))
}

func robMinu(n []int) int {
	a, b := n[0], max(n[0], n[1])
	for _, v := range n[2:] {
		a, b = b, max(a+v, b)
	}
	return max(a, b)
}

// 太贪了
// 想要一次遍历, 求最大值问题, 还记录是否包含 [0] 元素

// 退一步
// 因为首位互斥, 就分别求解包含首位的结果, 取最大值
