package leetcode

import "fmt"

func main() {
	fmt.Println("hello go for leetcode")

	arr := []int{4, 5, 2, 4, 3, 3, 1, 2, 5, 4}
	fee := 1
	test(arr, fee)
}

func test(prices []int, fee int) int {
	size := len(prices)
	if size < 2 {
		return 0
	}
	// res := 0
	var res, start, th int
	for i := 0; i < size; {
		fmt.Println("s: ", i)
		start = prices[i]
		th = start + fee
		for ; i < size; i++ {
			if th-prices[i] > fee {
				break
			}
			if start > prices[i] {
				fmt.Println("s: ", i)
				start = prices[i]
				th = start + fee
			} else if th < prices[i] {
				th = prices[i]
			}
		}

		if th-start > fee {
			res += th - start - fee
		}
		fmt.Println("e: ", i)
		fmt.Println("res: ", res)
	}
	return res
}
