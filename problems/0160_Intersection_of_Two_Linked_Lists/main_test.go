package _0160

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// head must Not be nil
func tailOf(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}

func TestGetIntersectionNode(t *testing.T) {
	Convey("get intersection node", t, func() {
		//    4->1-|
		// 		   8->4->5
		// 5->0->1-|
		ha := util.IntsToSList([]int{4, 1})
		hb := util.IntsToSList([]int{5, 0, 1})
		tail := util.IntsToSList([]int{8, 4, 5})
		tailA := tailOf(ha)
		tailB := tailOf(hb)
		tailA.Next = tail
		tailB.Next = tail
		So(_getIntersectionNode(ha, hb), ShouldEqual, tail)
		// 0->9->1-|
		// 		   2->4
		//       3-|
		ha = util.IntsToSList([]int{0, 9, 1})
		hb = util.IntsToSList([]int{3})
		tail = util.IntsToSList([]int{2, 4})
		tailA = tailOf(ha)
		tailB = tailOf(hb)
		tailA.Next = tail
		tailB.Next = tail
		So(_getIntersectionNode(ha, hb), ShouldEqual, tail)
		// 2->6->4
		//
		//    1->5
		ha = util.IntsToSList([]int{2, 6, 4})
		hb = util.IntsToSList([]int{1, 5})
		// tail = util.IntsToSList([]int{})
		// tailA = tailOf(ha)
		// tailB = tailOf(hb)
		// tailA.Next = tail
		// tailB.Next = tail
		So(_getIntersectionNode(ha, hb), ShouldBeNil)
	})
}
