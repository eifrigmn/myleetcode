package _0155

type node struct {
	val int
	prv *node
}
type MinStack struct {
	top  *node
	sTop *node
	len  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	// list
	this.top = &node{x, this.top}
	// sorted list
	sl := &node{0, this.sTop}
	curr := sl
	for curr.prv != nil && curr.prv.val < x {
		curr = curr.prv
	}
	curr.prv = &node{x, curr.prv}
	this.sTop = sl.prv
	this.len++
}

func (this *MinStack) Pop() {
	if this.len == 0 {
		return
	}
	n := this.top
	this.top = n.prv
	if this.sTop.val == n.val {
		this.sTop = this.sTop.prv
		return
	}
	// 遍历有序链表
	var curr = this.sTop
	for curr != nil && curr.prv != nil && curr.prv.val != n.val {
		curr = curr.prv
	}

	curr.prv = curr.prv.prv
	this.len--
}

func (this *MinStack) Top() int {
	if this.len == 0 {
		return 0
	}
	return this.top.val
}

func (this *MinStack) GetMin() int {
	if this.len == 0 {
		return 0
	}
	return this.sTop.val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
