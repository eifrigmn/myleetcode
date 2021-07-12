package _02

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstBadVersion(t *testing.T) {
	Convey("第一个错误的版本", t, func() {
		// n = 5, bad = 4
		var n = 5
		bad = 4
		So(firstBadVersion(n), ShouldEqual, bad)
		// n = 1, bad = 1
		n = 1
		bad = 1
		So(firstBadVersion(n), ShouldEqual, bad)
	})
}
