package _04

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	Convey("containsDuplicate", t, func(){
		// [1,2,3,1]
		nums := []int{1,2,3,1}
		So(containsDuplicate(nums), ShouldBeTrue)
		// [1,2,3,4]
		nums = []int{1,2,3,4}
		So(containsDuplicate(nums), ShouldBeFalse)
		// [1,1,1,3,3,4,3,2,4,2]
		nums = []int{1,1,1,3,3,4,3,2,4,2}
		So(containsDuplicate(nums), ShouldBeTrue)
	})
}
