package _01

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClimbStairs(t *testing.T) {
	Convey("爬楼梯", t, func() {
		// 1	1
		var n = 1
		So(climbStairs(n), ShouldEqual, 1)
		// 2	2
		n = 2
		So(climbStairs(n), ShouldEqual, 2)
		// 3	3
		n = 3
		So(climbStairs(n), ShouldEqual, 3)
		// 4	5
		n = 4
		So(climbStairs(n), ShouldEqual, 5)
		// 5	8
		n = 5
		So(climbStairs(n), ShouldEqual, 8)
		// 6	13
		n = 6
		So(climbStairs_math(n), ShouldEqual, 13)
		r := climbStairs(44)
		t.Log(r)
	})
}
