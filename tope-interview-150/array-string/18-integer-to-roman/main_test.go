package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntToRoman(t *testing.T) {
	Convey("整数转罗马数字", t, func() {
		num := 3
		result := intToRoman(num)
		So(result, ShouldEqual, "III")

		num = 4
		result = intToRoman(num)
		So(result, ShouldEqual, "IV")

		num = 9
		result = intToRoman(num)
		So(result, ShouldEqual, "IX")

		num = 58
		result = intToRoman(num)
		So(result, ShouldEqual, "LVIII")

		num = 1994
		result = intToRoman(num)
		So(result, ShouldEqual, "MCMXCIV")
	})
}
