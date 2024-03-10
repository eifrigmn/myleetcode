package _0083

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPartitionList(t *testing.T) {
	Convey("partition list", t, func() {
		// 输入：1->4->3->2->5->2, x=3
		// 输出：1->2->2->4->3->5
		l1 := util.IntsToSList([]int{1, 4, 3, 2, 5, 2})
		newl1 := _partition(l1, 3)
		newl1Arry := util.SListToInts(newl1)
		So(reflect.DeepEqual(newl1Arry, []int{1, 2, 2, 4, 3, 5}), ShouldBeTrue)
		// 输入：2->1，x=2
		// 输出：1->2
		l1 = util.IntsToSList([]int{2, 1})
		newl1 = _partition(l1, 2)
		newl1Arry = util.SListToInts(newl1)
		So(reflect.DeepEqual(newl1Arry, []int{1, 2}), ShouldBeTrue)
	})
}
