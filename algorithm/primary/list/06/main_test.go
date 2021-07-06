package _06

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHasCycle(t *testing.T) {
	Convey("环形链表", t, func() {
		// 3->2->0->-4---|
		//	  ^			 |
		//    |__________|
		l1 := util.Ints2ListWithCycle([]int{3, 2, 0, -4}, 1)
		So(hasCycle2(l1), ShouldBeTrue)
		// 1->2-|
		// ^    |
		// |____|
		l2 := util.Ints2ListWithCycle([]int{1, 2}, 0)
		So(hasCycle1(l2), ShouldBeTrue)
		// 1
		l3 := util.Ints2ListWithCycle([]int{1}, -1)
		So(hasCycle1(l3), ShouldBeFalse)
		// 1->2
		l4 := util.Ints2ListWithCycle([]int{1, 2}, -1)
		So(hasCycle2(l4), ShouldBeFalse)
	})
}
