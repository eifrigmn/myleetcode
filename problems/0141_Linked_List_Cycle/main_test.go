package _0141

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHasCycle(t *testing.T) {
	Convey("list has cycle", t, func() {
		// 3->2->0->-4---|
		//	  ^			 |
		//    |__________|
		l1 := util.Ints2ListWithCycle([]int{3, 2, 0, -4}, 1)
		So(_hasCycle(l1), ShouldBeTrue)
		// 1->2-|
		// ^    |
		// |____|
		l2 := util.Ints2ListWithCycle([]int{1, 2}, 0)
		So(_hasCycle(l2), ShouldBeTrue)
		// 1
		l3 := util.Ints2ListWithCycle([]int{1}, -1)
		So(_hasCycle(l3), ShouldBeFalse)
	})
}
