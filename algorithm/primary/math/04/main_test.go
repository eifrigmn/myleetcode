package _04

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRomanToInt(t *testing.T) {
	Convey("罗马数字", t, func() {
		// "III" -> 3
		var s = "III"
		var result = 3
		So(romanToInt(s), ShouldEqual, result)
		// "IV" -> 4
		s = "IV"
		result = 4
		So(romanToInt(s), ShouldEqual, result)
		// "IX" -> 9
		s = "IX"
		result = 9
		So(romanToInt(s), ShouldEqual, result)
		// "LVIII" -> 58
		s = "LVIII"
		result = 58
		So(romanToInt(s), ShouldEqual, result)
		// "MCMXCIV" -> 1994
		s = "MCMXCIV"
		result = 1994
		So(romanToInt(s), ShouldEqual, result)
	})
}
