package util

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListToInts(t *testing.T) {
	Convey("list to ints", t, func() {
		// 1 -> 2 -> 3 -> 4
		l := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		}
		target := []int{1, 2, 3, 4}
		nums := SListToInts(l)
		So(reflect.DeepEqual(nums, target), ShouldBeTrue)
	})
}

func TestIntsToSList(t *testing.T) {
	Convey("ints to singly list", t, func() {
		// [1,2,3,4]
		nums := []int{1, 2, 3, 4}
		list := IntsToSList(nums)
		var i = 0
		head := list
		for head != nil {
			So(head.Val, ShouldEqual, nums[i])
			i++
			head = head.Next
		}
	})
}

// func TestInts2ListWithCycle(t *testing.T) {
// 	Convey("ints to circular linked list", t, func() {
// 		nums := []int{3, 2, 0, -4}
// 		pos := 1
// 		list := Ints2ListWithCycle(nums, pos)
// 		var i = 0
// 		head := list
// 		for head != nil {
// 			So(head.Val, ShouldEqual, nums[i])
// 			i++
// 			head = head.Next
// 		}
// 	})
// }
