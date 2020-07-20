package _03

import (
	. "github.com/smartystreets/goconvey/convey"
	"myleetcode/util"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	Convey("反转链表", t, func() {
		// 1->2->3->4->5->NULL
		var rl, l *ListNode
		l = util.IntsToSList([]int{1,2,3,4,5})
		//rl = reverseList(l)
		rl = reverseList1(l)
		So(reflect.DeepEqual(util.SListToInts(rl), []int{5,4,3,2,1}), ShouldBeTrue)
	})
}


