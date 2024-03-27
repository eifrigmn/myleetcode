package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRomanToInt(t *testing.T) {
	Convey("罗马数字转整形", t, func() {
		s := "III"
		result := romanToInt(s)
		So(result, ShouldEqual, 3)

		s = "IV"
		result = romanToInt(s)
		So(result, ShouldEqual, 4)

		s = "IX"
		result = romanToInt(s)
		So(result, ShouldEqual, 9)

		s = "LVIII"
		result = romanToInt(s)
		So(result, ShouldEqual, 58)

		s = "MCMXCIV"
		result = romanToInt(s)
		So(result, ShouldEqual, 1994)

	})
}
