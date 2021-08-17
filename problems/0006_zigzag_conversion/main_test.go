package _0006

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	Convey("convert", t, func() {
		s := "PAYPALISHIRING"
		numRows := 3
		result := convert(s, numRows)
		So(result, ShouldEqual, "PAHNAPLSIIGYIR")

		numRows = 4
		result = convert(s, numRows)
		So(result, ShouldEqual, "PINALSIGYAHRPI")

		s = "A"
		numRows = 1
		result = convert(s, numRows)
		So(result, ShouldEqual, "A")
	})
}
