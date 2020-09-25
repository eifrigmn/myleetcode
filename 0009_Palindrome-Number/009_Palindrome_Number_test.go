package _0009

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	Convey("isPalindrome", t, func() {
		data := 12321
		flag := isPalindrome(data)
		So(flag, ShouldBeTrue)
	})
}
