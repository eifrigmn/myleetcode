package _0005

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	Convey("longestPalindrome", t, func() {
		s := "babad"
		targets := []string{"bab", "aba"}
		result := longestPalindrome(s)
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
		result = longestPalindrome(s)
		So(result, ShouldEqual, target)
	})
}
