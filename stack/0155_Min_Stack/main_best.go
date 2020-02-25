package stack 

type node2 struct {
	val int
	prv *node
}
type MinStack2 struct {
	top  *node
	sTop *node
	len  int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{}
}

func (this *MinStack2) Push(x int) {
	// list
	this.top = &node{x, this.top}
	// sorted list
	if this.len == 0 {
		this.sTop = &node{x, nil}
	} else {
		if this.sTop.val > x {
			this.sTop = &node{x, this.sTop}
		} else {
			this.sTop = &node{this.sTop.val, this.sTop}
		}
	}
	this.len++
}

func (this *MinStack2) Pop() {
	if this.len == 0 {
		return
	}
	this.sTop = this.sTop.prv
	this.top = this.top.prv
	this.len--
}

func (this *MinStack2) Top() int {
	if this.len == 0 {
		return 0
	}
	return this.top.val
}

func (this *MinStack2) GetMin() int {
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
