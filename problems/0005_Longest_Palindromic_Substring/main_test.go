package _0005

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestPalindrome(t *testing.T) {
	Convey("longestPalindrome", t, func() {
		s := "babad"
		targets := []string{"bab", "aba"}
		result := _longestPalindrome(s)
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
		result = _longestPalindrome(s)
		So(result, ShouldEqual, target)

		s = "a"
		result = _longestPalindrome(s)
		So(result, ShouldEqual, "a")

		s = "ac"
		result = _longestPalindrome(s)
		So(result, ShouldEqual, "a")

		s = "aacabdkacaa"
		result = _longestPalindrome(s)
		So(result, ShouldEqual, "aca")
	})
}

func TestSlice(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5:5]
	t.Log("s1", len(s1), cap(s1))
	s1 = append(s1, 100)
	t.Log(s1)
	t.Log(slice)
	//s2 := s1[2:6:7]
	//t.Log(s2)
	//t.Log(len(s2))
	//t.Log(s2[4])
	//s2 = append(s2, 100)
	//t.Log(s2)
}
