package _002

import (
	. "github.com/smartystreets/goconvey/convey"
	"myleetcode/util"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	Convey("addTwoNumbers", t, func() {
		num1 := []int{2, 4, 3}
		num2 := []int{5, 6, 4}
		target := []int{7, 0, 8}
		l1 := util.IntsToSList(num1)
		l2 := util.IntsToSList(num2)
		result := addTwoNumbers(l1, l2)
		resultArr := util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)
		num1 = []int{5}
		num2 = []int{5}
		target = []int{0, 1}
		result = addTwoNumbers(util.IntsToSList(num1), util.IntsToSList(num2))
		resultArr = util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)

		num1 = []int{0}
		num2 = []int{7, 3}
		target = []int{7, 3}
		result = addTwoNumbers(util.IntsToSList(num1), util.IntsToSList(num2))
		resultArr = util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)
	})
}
