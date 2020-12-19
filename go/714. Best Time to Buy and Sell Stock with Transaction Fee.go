package leetcode

func maxProfit(prices []int, fee int) int {
	size := len(prices)
	if size < 2 {
		return 0
	}
	// res := 0
	var res, start, th int
	for i := 0; i < size; {
		start = prices[i]
		th = start + fee
		for ; i < size; i++ {
			if th-prices[i] > fee {
				break
			} else if start > prices[i] {
				start = prices[i]
				th = start + fee
			} else if th < prices[i] {
				th = prices[i]
			}
		}

		if th-start > fee {
			res += th - start - fee
		}
	}
	return res
}
func maxProfit2(prices []int, fee int) int {
	size := len(prices)
	profit := 0
	buy := prices[0] + fee
	for i := 1; i < size; i++ {
		if buy > prices[i]+fee {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
