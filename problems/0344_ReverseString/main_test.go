package _0350

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	Convey("反转字符串", t, func() {
		s := []byte("hello")
		result := []byte("olleh")
		reverseString(s)
		So(s, ShouldEqual, result)
	})
}
