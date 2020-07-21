package _04

import (
	. "github.com/smartystreets/goconvey/convey"
	"myleetcode/util"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	Convey("合并两个有序链表", t, func() {
		var l1, l2, l *ListNode
		// 1->2->3  1->2->4
		l1 = util.IntsToSList([]int{1,2,3})
		l2 = util.IntsToSList([]int{1,2,4})
		l = mergeTwoLists(l1, l2)
		So(reflect.DeepEqual(util.SListToInts(l), []int{1,1,2,2,3,4}), ShouldBeTrue)
		// 1->2->3   2->4->6->8
		l1 = util.IntsToSList([]int{1,2,3})
		l2 = util.IntsToSList([]int{2,4,6,8})
		l = mergeTwoLists(l1, l2)
		So(reflect.DeepEqual(util.SListToInts(l), []int{1,2,2,3,4,6,8}), ShouldBeTrue)
		// 1->2->4, 1->3->4
		l1 = util.IntsToSList([]int{1,2,4})
		l2 = util.IntsToSList([]int{1,3,4})
		l = mergeTwoLists(l1, l2)
		So(reflect.DeepEqual(util.SListToInts(l), []int{1,1,2,3,4,4}), ShouldBeTrue)
		// 2, 1
		l1 = util.IntsToSList([]int{2})
		l2 = util.IntsToSList([]int{1})
		l = mergeTwoLists(l1, l2)
		So(reflect.DeepEqual(util.SListToInts(l),[]int{1,2}), ShouldBeTrue)
		// 5, 1->2->4
		l1 = util.IntsToSList([]int{5})
		l2 = util.IntsToSList([]int{1,2,4})
		l = mergeTwoLists(l1, l2)
		So(reflect.DeepEqual(util.SListToInts(l), []int{1,2,4,5}), ShouldBeTrue)
	})
}
