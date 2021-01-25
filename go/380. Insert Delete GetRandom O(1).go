package leetcode

import "math/rand"

// RandomizedSet 随机set集合
type RandomizedSet struct {
	dict map[int]int
	arr  []int
}

// Constructor Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		make(map[int]int), []int{},
	}
}

// Insert Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {
	if _, has := rs.dict[val]; has {
		return false
	}
	rs.dict[val] = len(rs.arr)
	rs.arr = append(rs.arr, val)
	return true
}

// Remove Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	if idx, has := rs.dict[val]; has {
		last := len(rs.arr) - 1
		rs.arr[idx], rs.arr[last] = rs.arr[last], rs.arr[idx]
		rs.dict[rs.arr[idx]] = idx
		delete(rs.dict, val)
		rs.arr = rs.arr[:last]
		return true
	}
	return false
}

// GetRandom Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	return rs.arr[rand.Intn(len(rs.arr))]
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
