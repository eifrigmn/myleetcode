package _01

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHammingWeight(t *testing.T) {
	Convey("位1的个数", t, func() {
		var num uint32
		num = 00000000000000000000000000001011
		So(hammingWeight(num), ShouldEqual, 3)
		num = 00000000000000000000000010000000
		So(hammingWeight(num), ShouldEqual, 1)
		// num = 11111111111111111111111111111101
		// So(hammingWeight(num), ShouldEqual, 31)
	})
}
