package _08

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	Convey("外观数列", t, func() {
		var n int
		// 1 -> "1"
		n= 1
		So(countAndSay(n), ShouldEqual, "1")
		// 2 -> "11"
		n = 2
		So(countAndSay(n), ShouldEqual, "11")
		// 3 -> "21"
		n = 3
		So(countAndSay(n), ShouldEqual, "21")
		// 4 -> "1211"
		n = 4
		So(countAndSay(n), ShouldEqual, "1211")
		// 4 -> "111221"
		n = 5
		So(countAndSay(n), ShouldEqual, "111221")
	})
}
