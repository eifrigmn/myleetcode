package _06

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	Convey("两个数组的交集 II", t, func(){
		nums1 := []int{1,2,2,1}
		nums2 := []int{2,2}
		result := intersect1(nums1, nums2)
		So(reflect.DeepEqual(result, []int{2,2}), ShouldBeTrue)
		nums1 = []int{4,9,5}
		nums2 = []int{9,4,9,8,4}
		result = intersect(nums1, nums2)
		So(reflect.DeepEqual(result, []int{4,9}), ShouldBeTrue)
	})
}