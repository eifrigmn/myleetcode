package _0005

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestPalindrome(t *testing.T) {
	Convey("longestPalindrome", t, func() {
		s := "babad"
		result := _longestPalindrome(s)
		So(result, ShouldBeIn, []string{"bab", "aba"})

		s = "cbbd"
		target := "bb"
		result = _longestPalindrome(s)
		So(result, ShouldEqual, target)

		s = "a"
		result = _longestPalindrome(s)
		So(result, ShouldEqual, "a")

		s = "ac"
		result = _longestPalindrome(s)
		So(result, ShouldBeIn, []string{"a", "c"})

		s = "aacabdkacaa"
		result = _longestPalindrome(s)
		So(result, ShouldEqual, "aca")
	})
}
