package _0142

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDetectCycle(t *testing.T) {
	Convey("环形链表II", t, func() {
		// 3->2->0->-4->5--|
		//	  ^			   |
		//    |____________|
		l1 := util.Ints2ListWithCycle([]int{3, 2, 0, -4, 5}, 1)
		node1 := _detectCycle(l1)
		So(node1, ShouldEqual, l1.Next)
		// 1->2--|
		// ^     |
		// |_____|
		l2 := util.Ints2ListWithCycle([]int{1, 2}, 0)
		node2 := _detectCycle(l2)
		So(node2, ShouldEqual, l2)
		// 1
		l3 := util.Ints2ListWithCycle([]int{1}, -1)
		node3 := _detectCycle(l3)
		So(node3, ShouldBeNil)
		// -1 -> -7 -> 7 -> -4 -> 19 -> 6 -> -9 -> -5 -> -2 -> -5 --|
		//	                                  ^			            |
		//                                    |_____________________|
		l4 := util.Ints2ListWithCycle([]int{-1, -7, 7, -4, 19, 6, -9, -5, -2, -5}, 6)
		node4 := _detectCycle(l4)
		target4 := l4
		for i := 0; i < 6; i++ {
			target4 = target4.Next
		}
		So(node4, ShouldEqual, target4)
	})
}
