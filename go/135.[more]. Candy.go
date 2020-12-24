package leetcode

// ac
/**
解题思路:
归纳法, 从 n-1 推广的 n, 考虑 n 的各种情况, 分别处理
*/

func candy(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return n
	}

	arr := make([]int, n)
	// pre := 1
	for i, v := range ratings {
		if i == 0 || v == ratings[i-1] {
			arr[i] = 1
		} else if v > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = 1
			for j := i - 1; j >= 0 && ratings[j] > ratings[j+1] && arr[j] <= arr[j+1]; j-- {
				arr[j] = arr[j+1] + 1
			}
		}
	}
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}

// todo	左右规则分别处理
// todo	信息记录, 根据增减规则处理 [这个没怎么读懂, 可能需要再捋捋, 住院, 就先搁着吧...]
