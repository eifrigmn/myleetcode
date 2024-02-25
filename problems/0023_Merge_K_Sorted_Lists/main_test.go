package _0021

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoLists(t *testing.T) {
	Convey("merge k sorted lists", t, func() {
		// 输入：
		//[
		//  1->4->5,
		//  1->3->4,
		//  2->6
		//]
		// 输出：1->1->2->3->4->4->5->6
		nums1 := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
		lists := util.IntsToSLists(nums1)
		result := mergeKLists(lists)
		resultArry := util.SListToInts(result)
		So(reflect.DeepEqual(resultArry, []int{1, 1, 2, 3, 4, 4, 5, 6}), ShouldBeTrue)
	})
}
