package _0019

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveNthFromEnd(t *testing.T) {
	Convey("删除链表倒数第N个节点", t, func() {
		// 1->2->3->4->5, n=2 ===> 1->2->3->5
		list := util.IntsToSList([]int{1, 2, 3, 4, 5})
		result := util.SListToInts(_removeNthFromEnd(list, 2))
		So(reflect.DeepEqual(result, []int{1, 2, 3, 5}), ShouldBeTrue)
		// 1->2->3->4->5, n=5 ===> 2->3->4->5
		list = util.IntsToSList([]int{1, 2, 3, 4, 5})
		result = util.SListToInts(_removeNthFromEnd(list, 5))
		So(reflect.DeepEqual(result, []int{2, 3, 4, 5}), ShouldBeTrue)
	})
}
