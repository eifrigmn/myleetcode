package linked_list

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveElements(t *testing.T) {
	Convey("remove linked list elements", t, func() {
		// 1->2->6->3->4->5->6
		l1 := util.IntsToSList([]int{1, 2, 6, 3, 4, 5, 6})
		rl1 := removeElements(l1, 6)
		So(reflect.DeepEqual(util.SListToInts(rl1), []int{1, 2, 3, 4, 5}), ShouldBeTrue)
		// 1->2->3->4
		l2 := util.IntsToSList([]int{1, 2, 3, 4})
		rl2 := removeElements(l2, 6)
		So(reflect.DeepEqual(util.SListToInts(rl2), []int{1, 2, 3, 4}), ShouldBeTrue)
		// 6
		l3 := util.IntsToSList([]int{6})
		rl3 := removeElements(l3, 6)
		var emptArry []int
		So(reflect.DeepEqual(util.SListToInts(rl3), emptArry), ShouldBeTrue)
		l4 := util.IntsToSList([]int{1, 1, 1})
		rl4 := removeElements(l4, 1)
		So(reflect.DeepEqual(util.SListToInts(rl4), emptArry), ShouldBeTrue)
	})
}
