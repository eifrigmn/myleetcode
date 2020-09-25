package _0148

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortList(t *testing.T) {
	Convey("链表排序", t, func() {
		// 4->2->1->3
		l := util.IntsToSList([]int{4, 2, 1, 3})
		sl := sortList(l)
		So(reflect.DeepEqual(util.SListToInts(sl), []int{1, 2, 3, 4}), ShouldBeTrue)
		// -1->5->3->4->0
		l = util.IntsToSList([]int{-1, 5, 3, 4, 0})
		sl = sortList(l)
		So(reflect.DeepEqual(util.SListToInts(sl), []int{-1, 0, 3, 4, 5}), ShouldBeTrue)
	})
}
