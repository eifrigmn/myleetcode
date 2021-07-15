package _03

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPowerOfThree(t *testing.T) {
	Convey("3的幂", t, func() {
		// 27 -> true
		So(isPowerOfThree(27), ShouldBeTrue)
		// 0 -> false
		So(isPowerOfThree(0), ShouldBeFalse)
		// 45 -> false
		So(isPowerOfThree(45), ShouldBeFalse)
	})
}
