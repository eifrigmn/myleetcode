package _011

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_maxArea(t *testing.T) {
	Convey("isPalindrome", t, func() {
		height := []int{1,8,6,2,5,4,8,3,7}
		So( maxArea(height), ShouldEqual, 49)
		height = []int{1,81}
		So( maxArea(height), ShouldEqual, 1)
	})
}
