package _04

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	Convey("有效的字母异位词", t, func(){
		s := "anagram"
		t := "nagaram"
		So(isAnagram(s, t), ShouldBeTrue)

		s = "rat"
		t = "car"
		So(isAnagram(s, t), ShouldBeFalse)
	})
}
