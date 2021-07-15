package _02

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountPrimes(t *testing.T) {
	Convey("计数质数", t, func() {
		// 10 -> 4
		So(countPrimes(10), ShouldEqual, 4)
		// 1 -> 0
		So(countPrimes(1), ShouldEqual, 0)
		// 6 -> 3
		So(countPrimes(6), ShouldEqual, 3)
		// 5 -> 3
		So(countPrimes(5), ShouldEqual, 2)
		// 2 -> 0
	})
}
