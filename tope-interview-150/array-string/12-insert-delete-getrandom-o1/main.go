package main

import "math/rand"

type RandomizedSet struct {
	content map[int]int
	nums    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		content: make(map[int]int),
		nums:    []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.content[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.content[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.content[val]
	if !ok {
		return false
	}
	lastIdx := len(this.nums) - 1
	// 把最后一位的值赋给在数组中被删除的值
	this.nums[idx] = this.nums[lastIdx]
	// map中记录调换后的值与游标的对应关系
	this.content[this.nums[idx]] = idx
	// 舍弃掉数组的最后一位
	this.nums = this.nums[:lastIdx]
	// 删除map中值为val的key
	delete(this.content, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
