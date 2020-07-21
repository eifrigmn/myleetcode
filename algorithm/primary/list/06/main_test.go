package _06

import (
	. "github.com/smartystreets/goconvey/convey"
	"myleetcode/util"
	"testing"
)

func TestHasCycle(t *testing.T) {
	Convey("环形链表", t, func() {
		// 3->2->0->-4---|
		//	  ^			 |
		//    |__________|
		l1 := util.Ints2ListWithCycle([]int{3, 2, 0, -4}, 1)
		So(hasCycle(l1), ShouldBeTrue)
		// 1->2-|
		// ^    |
		// |____|
		l2 := util.Ints2ListWithCycle([]int{1,2}, 0)
		So(hasCycle(l2), ShouldBeTrue)
		// 1
		l3 := util.Ints2ListWithCycle([]int{1}, -1)
		So(hasCycle(l3), ShouldBeFalse)
	})
}
