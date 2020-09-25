package _0143

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReorderList(t *testing.T) {
	Convey("重排链表", t, func() {
		// 1->2->3->4 ==> 1->4->2->3
		l1 := util.IntsToSList([]int{1, 2, 3, 4})
		reorderList(l1)
		So(reflect.DeepEqual(util.SListToInts(l1), []int{1, 4, 2, 3}), ShouldBeTrue)
		// 1->2->3->4->5 ==> 1->5->2->4->3
		l2 := util.IntsToSList([]int{1, 2, 3, 4, 5})
		reorderList(l2)
		So(reflect.DeepEqual(util.SListToInts(l2), []int{1, 5, 2, 4, 3}), ShouldBeTrue)

	})
}
