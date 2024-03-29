package _0206

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseList(t *testing.T) {
	Convey("reverse list", t, func() {
		// 1->2->3->4->5->NULL
		head := util.IntsToSList([]int{1, 2, 3, 4, 5})
		rHead := _reverseList(head)
		So(reflect.DeepEqual(util.SListToInts(rHead), []int{5, 4, 3, 2, 1}), ShouldBeTrue)
	})
}
