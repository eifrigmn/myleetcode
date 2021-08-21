package _0009

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPalindrome(t *testing.T) {
	Convey("isPalindrome", t, func() {
		data := 12321
		flag := isPalindrome2(data)
		So(flag, ShouldBeTrue)
	})
}
