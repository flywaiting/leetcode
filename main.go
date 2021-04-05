package main

import (
	"fmt"
	"leetcode/glc"
)

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	glc.Merge(a, b, 3, 3)
	fmt.Println(a)
}
