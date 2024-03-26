package main

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("除自身以外数组的乘积", t, func() {
		nums := []int{1, 2, 3, 4}
		result := productExceptSelf(nums)
		So(reflect.DeepEqual(result, []int{24, 12, 8, 6}), ShouldBeTrue)

		nums = []int{-1, 1, 0, -3, 3}
		result = productExceptSelf(nums)
		So(reflect.DeepEqual(result, []int{0, 0, 9, 0, 0}), ShouldBeTrue)

	})
}
