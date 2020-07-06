package _01

import (
	"myleetcode/algorithm/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseString(t *testing.T) {
	Convey("反转字符串", t, func() {
		// ["h", "e", "l", "l", "o"]
		// ["o", "l", "l", "e", "h"]
		b := util.StrArry2BytesArry([]string{"h", "e", "l", "l", "o"})
		reverseString(b)
		s := util.BytesArray2StrArray(b)
		So(reflect.DeepEqual(s, []string{"o", "l", "l", "e", "h"}), ShouldBeTrue)

		// ["H","a","n","n","a","h"]
		// ["h","a","n","n","a","H"]
		b = util.StrArry2BytesArry([]string{"H", "a", "i", "n", "a", "h"})
		reverseString(b)
		s = util.BytesArray2StrArray(b)
		So(reflect.DeepEqual(s, []string{"h", "a", "n", "i", "a", "H"}), ShouldBeTrue)

	})
}
