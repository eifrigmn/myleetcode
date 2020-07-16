package _06

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	Convey("字符串转换整数 (atoi)", t, func() {
		var str string
		// "42"
		str = "42"
		So(myAtoi(str), ShouldEqual, 42)
		// "   -42"
		str = "   -42"
		So(myAtoi(str), ShouldEqual, -42)
		// "4193 with words"
		str = "4193 with words"
		So(myAtoi(str), ShouldEqual, 4193)
		// "words and 987"
		str = "words and 987"
		So(myAtoi(str), ShouldEqual, 0)
		// "-91283472332"
		str = "-91283472332"
		So(myAtoi(str), ShouldEqual, -2147483648)
		// "4193 with words 789"
		str = "4193 with words 789"
		So(myAtoi(str), ShouldEqual, 4193)
		// "+1"
		str = "+1"
		So(myAtoi(str), ShouldEqual, 1)
		// "+-2"
		str = "+-2"
		So(myAtoi(str), ShouldEqual, 0)
		// "9223372036854775808"
		str = "9223372036854775808"
		So(myAtoi(str), ShouldEqual, 2147483647)
	})
}
