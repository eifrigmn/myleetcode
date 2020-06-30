package _07

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	Convey("加一", t, func() {
		nums := []int{9, 9}
		result := plusOne(nums)
		So(reflect.DeepEqual(result, []int{1, 0, 0}), ShouldBeTrue)
		nums = []int{8,9,9}
		result = plusOne(nums)
		So(reflect.DeepEqual(result, []int{9,0,0}), ShouldBeTrue)
		nums = []int{1,2,3}
		result = plusOne(nums)
		So(reflect.DeepEqual(result, []int{1,2,4}), ShouldBeTrue)
		nums = []int{4,3,2,1}
		result = plusOne(nums)
		So(reflect.DeepEqual(result, []int{4,3,2,2}), ShouldBeTrue)
	})
}
