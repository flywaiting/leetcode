package leetcode

import "math/rand"

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
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

//  想要实现常数时间的插入, 删除操作, 哈希表是首选

// !	learn:
// golang 的 map 没有随机方法, 也没办法直接获取 key/val 列表, 需要额外维护
// 维护 slice, 方便获取随机值
// 维护 map, 映射 slice 的索引
// 让 slice 作为类 stack 使用, 提高 remove 操作的性能, 【将待删除值与切片末尾交换, 然后去除末尾】
// notice 交换之后, map 对应值要一起改变, 不然就GG
