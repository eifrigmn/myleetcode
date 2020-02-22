package linked_list

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoLists(t *testing.T) {
	Convey("merge two lists", t, func() {
		// 1->2->4
		nums1 := []int{1, 2, 4}
		l1 := util.IntsToSList(nums1)
		// 1->3->4
		nums2 := []int{1, 3, 4}
		l2 := util.IntsToSList(nums2)
		result := mergeTwoLists(l1, l2)
		resultArry := util.SListToInts(result)
		So(reflect.DeepEqual(resultArry, []int{1, 1, 2, 3, 4, 4}), ShouldBeTrue)
	})
}
