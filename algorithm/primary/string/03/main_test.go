package _03

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	Convey("字符串中的第一个唯一字符", t, func(){
		s := "leetcode"
		index := firstUniqChar(s)
		So(index, ShouldEqual, 0)

		s = "loveleetcode"
		index = firstUniqChar(s)
		So(index, ShouldEqual, 2)

		s = "aabbccdd"
		index = firstUniqChar(s)
		So(index, ShouldEqual, -1)
	})
}
