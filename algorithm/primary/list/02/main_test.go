package _02

import (
	. "github.com/smartystreets/goconvey/convey"
	"myleetcode/util"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	Convey("删除链表的倒数第N个节点", t, func() {
		// 1->2->3->4->5, n = 2
		var l *ListNode
		var n int
		var result []int
		l = util.IntsToSList([]int{1,2,3,4,5})
		n = 2
		l = removeNthFromEnd(l, n)
		result = util.SListToInts(l)
		So(reflect.DeepEqual(result, []int{1,2,3,5}), ShouldBeTrue)
		// 1->2->3->4->5, n = 5
		l = util.IntsToSList([]int{1,2,3,4,5})
		n = 5
		l = removeNthFromEnd(l, n)
		result = util.SListToInts(l)
		So(reflect.DeepEqual(result, []int{2,3,4,5}), ShouldBeTrue)
		// 1->2->3->4->5, n = 1
		l = util.IntsToSList([]int{1,2,3,4,5})
		n = 1
		l = removeNthFromEnd(l, n)
		result = util.SListToInts(l)
		So(reflect.DeepEqual(result, []int{1,2,3,4}), ShouldBeTrue)
		// 1->2, n=1
		l = util.IntsToSList([]int{1,2})
		n = 1
		l = removeNthFromEnd(l, n)
		result = util.SListToInts(l)
		So(reflect.DeepEqual(result, []int{1}), ShouldBeTrue)
	})
}
