package _02

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReverse(t *testing.T) {
	Convey("整数反转", t, func(){
		// 123 -> 321
		num := 123
		rnum := reverse(num)
		So(rnum, ShouldEqual, 321)

		// -123 -> -321
		num = -123
		rnum = reverse(num)
		So(rnum, ShouldEqual, -321)

		// 120 -> 21
		num = 120
		rnum = reverse(num)
		So(rnum, ShouldEqual, 21)

		// 1534236469 -> 0
		num = 1534236469
		rnum = reverse(num)
		So(rnum, ShouldEqual, 0)
	})
}
