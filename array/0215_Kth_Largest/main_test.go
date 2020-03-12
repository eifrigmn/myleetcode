package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindKthLargest(t *testing.T) {
	Convey("find kth largest", t, func() {
		nums1 := []int{3, 2, 1, 5, 6, 4}
		k1 := 2
		So(findKthLargest2(nums1, k1), ShouldEqual, 5)
		nums2 := []int{3,2,3,1,2,4,5,5,6}
		k2 := 4
		So(findKthLargest2(nums2, k2), ShouldEqual, 4)
	})
}

// 参考
type heap struct {
	cap  int
	data []int
}

func construct(k int) *heap {
	h := &heap{}
	h.cap = k
	h.data = []int{}
	return h
}

func (this *heap) Push(ele int) {
	if len(this.data) < this.cap {
		this.data = append(this.data, ele)
		this.up(len(this.data) - 1)
		return
	}
	if ele > this.data[0] {
		this.data[0] = ele
		this.down(0)
		return
	}
}

func (this *heap) up(i int) {
	for i != 0 {
		j := (i+1)/2 - 1
		if this.data[i] < this.data[j] {
			this.data[i], this.data[j] = this.data[j], this.data[i]
		} else {
			break
		}
		i = j
	}
}

func (this *heap) down(i int) {
	for i != len(this.data)-1 && 2*i+1 < len(this.data) {
		j := 2*i + 1
		if j+1 < len(this.data) && this.data[2*i+1] > this.data[2*i+2] {
			j = 2*i + 2
		}
		if this.data[i] > this.data[j] {
			this.data[i], this.data[j] = this.data[j], this.data[i]
		} else {
			break
		}
		i = j
	}
}

func findKthLargest1(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	h := construct(k)
	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
	}
	return h.data[0]
}
