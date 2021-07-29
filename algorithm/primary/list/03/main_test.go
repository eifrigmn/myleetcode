package _03

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseList(t *testing.T) {
	Convey("反转链表", t, func() {
		// 1->2->3->4->5->NULL
		var rl, l *ListNode
		l = util.IntsToSList([]int{1, 2, 3, 4, 5})
		//rl = reverseList(l)
		rl = reverseList_stack(l)
		So(reflect.DeepEqual(util.SListToInts(rl), []int{5, 4, 3, 2, 1}), ShouldBeTrue)
	})
}
