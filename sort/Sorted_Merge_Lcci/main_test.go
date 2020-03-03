package sort 

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)
func TestMerge(t *testing.T) {
	Convey("合并两个排序数组", t, func(){
		// A = [1,2,3,0,0,0], m = 3
		// B = [2,5,6],       n = 3
		A := []int{1,2,3,0,0,0}
		B := []int{2,5,6}
		m := 3
		n := 3
		merge(A, m, B, n)
		So(reflect.DeepEqual(A, []int{1,2,2,3,5,6}), ShouldBeTrue)
		// A = [0], m = 0
		// B = [1], n = 1
		A = []int{0}
		B = []int{1}
		m = 0
		n = 1
		merge(A, m, B, n)
		So(reflect.DeepEqual(A, []int{1}), ShouldBeTrue)
		// A = [2,0], m = 1
		// B = [1],   n = 1
		A = []int{2,0}
		B = []int{1}
		m = 1
		n = 1
		merge(A, m, B, n)
		So(reflect.DeepEqual(A, []int{1,2}), ShouldBeTrue)
	})
}