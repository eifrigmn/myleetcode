package stack

import "sort"

type MinStack1 struct {
	items       []int
	sortedItems []int
	len         int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{}
}

func (this *MinStack1) Push(x int) {
	this.items = append(this.items, x)
	this.sortedItems = append(this.sortedItems, x)
	sort.Ints(this.sortedItems)
	this.len++
}

func (this *MinStack1) Pop() {
	if this.len == 0 {
		return
	}

	var i int
	for i < this.len && this.sortedItems[i] != this.items[this.len-1] {
		i++
	}

	if i == this.len-1 {
		this.sortedItems = this.sortedItems[:this.len-1]
	} else {
		this.sortedItems = append(this.sortedItems[:i], this.sortedItems[i+1:]...)
	}
	this.items = this.items[:this.len-1]
	this.len--
}

func (this *MinStack1) Top() int {
	return this.items[this.len-1]
}

func (this *MinStack1) GetMin() int {
	if this.len == 0 {
		return 0
	}
	return this.sortedItems[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
