package main

import (
	"fmt"
	"leetcode/glc"
)

func main() {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans := glc.Trap(h)
	fmt.Println(ans)

}
