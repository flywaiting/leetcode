package leetcode

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 简单的递归

// 还有直接的数学公式、循环
