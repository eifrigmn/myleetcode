package _016

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	Convey("isPalindrome", t, func() {
		So( threeSumClosest([]int{-1,2,1,-4}, 1), ShouldEqual, 2)
		So( threeSumClosest([]int{0,0,0}, 1), ShouldEqual, 0)
	})
}
