package _0075

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortColors(t *testing.T) {
	Convey("颜色分类", t, func(){
		// [2,0,2,1,1,0] -> [0,0,1,1,2,2]
		nums := []int{2,0,2,1,1,0}
		sortedNums := []int{0,0,1,1,2,2}
		sortColors1(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}