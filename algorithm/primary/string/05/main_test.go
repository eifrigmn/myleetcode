package _05

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	Convey("验证回文字符串", t, func() {
		var s string
		// "A man, a plan, a canal: Panama"
		s = "A man, a plan, a canal: Panama"
		So(isPalindrome(s), ShouldBeTrue)
		// "race a car"
		s = "race a car"
		So(isPalindrome(s), ShouldBeFalse)
	})
}
