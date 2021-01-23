package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("hello go for leetcode")

	rs := Constructor()
	fmt.Println("i:", rs.Insert(0))
	fmt.Println("i:", rs.Insert(1))
	fmt.Println("r:", rs.Remove(0))
	fmt.Println("i:", rs.Insert(2))
	fmt.Println("r:", rs.Remove(1))
	fmt.Println(rs.GetRandom())
}

type RandomizedSet struct {
	dict map[int]int
	arr  []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		make(map[int]int), []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, has := this.dict[val]; has {
		return false
	}
	this.dict[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, has := this.dict[val]; has {
		last := len(this.arr) - 1
		this.arr[idx], this.arr[last] = this.arr[last], this.arr[idx]
		this.dict[this.arr[idx]] = idx
		delete(this.dict, val)
		this.arr = this.arr[:last]

		fmt.Println(val)
		fmt.Println(this.dict)
		fmt.Println(this.arr)

		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	fmt.Println(this.arr)
	return this.arr[rand.Intn(len(this.arr))]
}
