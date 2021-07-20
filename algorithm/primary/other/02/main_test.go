package _02

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHammingDistance(t *testing.T) {
	Convey("汉明距离", t, func() {
		// x=1,y=4 -> 2
		So(hammingDistance(1, 4), ShouldEqual, 2)
		// x=3,y=1 -> 1
		So(hammingDistance(3, 1), ShouldEqual, 1)
	})
}
