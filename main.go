package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello go for leetcode")

	s := []int{1, 2, 0, 0}
	// canPlaceFlowers(s, 2)

	addToArrayForm(s, 9000)

	fmt.Println([]int{123})
}

func addToArrayForm(A []int, K int) []int {
	if K == 0 {
		return A
	}
	tmp := 0
	for _, v := range A {
		tmp = v + tmp*10
	}
	fmt.Println(tmp)
	tmp += K

	ans := make([]int, 0, len(A))
	for ; tmp > 0; tmp /= 10 {
		ans = append(ans, tmp%10)
		fmt.Println(ans, tmp)
	}
	return ans
}
