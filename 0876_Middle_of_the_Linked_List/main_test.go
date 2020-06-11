package _0876

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMiddleNode(t *testing.T) {
	Convey("链表的中点", t, func() {
		// 1->2->3->4->5
		lft := util.IntsToSList([]int{1, 2})
		mid := util.IntsToSList([]int{3, 4, 5})
		head := lft
		lftTail := util.TailOf(lft)
		lftTail.Next = mid
		So(middleNode(head), ShouldEqual, mid)
		// 1->2->3->4->5->6
		lft = util.IntsToSList([]int{1, 2, 3})
		mid = util.IntsToSList([]int{4, 5, 6})
		head = lft 
		lftTail = util.TailOf(lft)
		lftTail.Next = mid
		So(middleNode(head), ShouldEqual, mid)
	})
}
