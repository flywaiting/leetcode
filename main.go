package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello go for leetcode")

	// s := []int{1, 0, 0, 0, 0}
	// canPlaceFlowers(s, 2)

	i, j := 10, i-2
	fmt.Println(i, j)
}
func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	if 2*n > size+1 {
		return false
	}

	pre, cnt := -1, 0
	for i, v := range flowerbed {
		if v == 1 {
			cnt += (i - pre - 2 + 1) / 2
			pre = i + 1
			fmt.Println(cnt, i)
		}
	}

	// cnt += (size - pre - 2 + 1) / 2
	// return cnt >= n
	fmt.Println(cnt)
	return cnt >= n
}
