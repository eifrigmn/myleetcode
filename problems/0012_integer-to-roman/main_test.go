package _012

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	Convey("isPalindrome", t, func() {
		So( intToRoman(3), ShouldEqual, "III")
		So( intToRoman(4), ShouldEqual, "IV")
		So( intToRoman(9), ShouldEqual, "IX")
		So( intToRoman(58), ShouldEqual, "LVIII")
		So( intToRoman(1994), ShouldEqual, "MCMXCIV")
		So( intToRoman(10), ShouldEqual, "X")
	})
}
