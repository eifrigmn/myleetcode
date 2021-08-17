package _0005

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestPalindrome(t *testing.T) {
	Convey("longestPalindrome", t, func() {
		s := "babad"
		targets := []string{"bab", "aba"}
		result := longestPalindrome2(s)
		var flag bool
		for _, target := range targets {
			if result == target {
				flag = true
				break
			}
		}
		So(flag, ShouldBeTrue)

		s = "cbbd"
		target := "bb"
		result = longestPalindrome2(s)
		So(result, ShouldEqual, target)

		s = "a"
		result = longestPalindrome2(s)
		So(result, ShouldEqual, "a")

		s = "ac"
		result = longestPalindrome2(s)
		So(result, ShouldEqual, "a")

	})
}
