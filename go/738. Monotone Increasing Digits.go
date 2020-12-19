package leetcode

func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}

	temp := N
	arr := []int{}
	for temp > 0 {
		arr = append(arr, temp%10)
		temp /= 10
	}

	n := len(arr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	flagArr := make([]bool, 0, n)
	flag := 0
	for i := 1; i < n; i++ {
		if flag > 0 {
			arr[i] = 9
			continue
		}
		if arr[i] < arr[i-1] {
			flag = 1
			arr[i] = 9
			for j := i - 1; i >= 0; j-- {
				if flagArr[j] {
					arr[j] = 9
				} else {
					arr[j]--
					break
				}
			}
		} else {
			flagArr[i] = arr[i] == arr[i-1]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = res*10 + arr[i]
	}
	return res
}

func monotoneIncreasingDigits2(N int) int {
	if N < 10 {
		return N
	}

	arr := make([]int, 0, 10)
	for N > 0 {
		arr = append(arr, N%10)
		N /= 10
	}

	size := len(arr)
	idx := size - 1
	for i := size - 2; i >= 0; i-- {
		if arr[i+1] <= arr[i] {
			idx--
		} else {
			break
		}
	}

	if idx > 0 {
		for ; idx < size; idx++ {
			if idx == size-1 || arr[idx] > arr[idx+1] {
				arr[idx]--
				idx--
				break
			}
		}
		for ; idx >= 0; idx-- {
			arr[idx] = 9
		}
	}
	res := 0
	size--
	for ; size >= 0; size-- {
		res = res*10 + arr[size]
	}
	return res
}

/**
解题思路：
从高位起，寻找第一个不满足递增的位置，然后向前回溯，找到第一个可以减少的位置【针对相等情况】，该位置后面的低位数全部替换最大值

ac之后，再整理思路，其实可以直接处理数据，不需要做翻转处理，只是人的习惯都是左高右低，
他人解答，多是转换成字符数组，再转换回来。怎么讲呢，不是很喜欢这种方法吧
*/
