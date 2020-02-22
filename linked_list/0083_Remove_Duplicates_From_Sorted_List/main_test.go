package linked_list

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeleteDuplicates(t *testing.T) {
	Convey("delete duplicate elements", t, func() {
		// 1->1->2
		l1 := util.IntsToSList([]int{1, 1, 2})
		newl1 := deleteDuplicates(l1)
		newl1Arry := util.SListToInts(newl1)
		So(reflect.DeepEqual(newl1Arry, []int{1, 2}), ShouldBeTrue)
	})
}
