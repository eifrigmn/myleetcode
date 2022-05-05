package _035

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	Convey("next permutation", t, func() {
		So(searchInsert([]int{1, 3, 5, 6}, 5), ShouldEqual, 2)
		So(searchInsert([]int{1, 3, 5, 6}, 2), ShouldEqual, 1)
		So(searchInsert([]int{1, 3, 5, 6}, 7), ShouldEqual, 4)
		So(searchInsert([]int{}, 5), ShouldEqual, 0)
		So(searchInsert([]int{1, 3, 5, 6}, 0), ShouldEqual, 0)
	})
}
