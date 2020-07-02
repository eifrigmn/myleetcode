package _09

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	Convey("两数之和", t, func() {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := twoSum(nums, target)
		So(reflect.DeepEqual(result, []int{0,1}), ShouldBeTrue)
	})
}
