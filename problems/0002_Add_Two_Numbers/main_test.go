package _002

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNumbers(t *testing.T) {
	Convey("addTwoNumbers", t, func() {
		num1 := []int{2, 4, 3}
		num2 := []int{5, 6, 4}
		target := []int{7, 0, 8}
		l1 := util.IntsToSList(num1)
		l2 := util.IntsToSList(num2)
		result := addTwoNumbers1(l1, l2)
		resultArr := util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)
		num1 = []int{5}
		num2 = []int{5}
		target = []int{0, 1}
		result = addTwoNumbers1(util.IntsToSList(num1), util.IntsToSList(num2))
		resultArr = util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)

		num1 = []int{0}
		num2 = []int{7, 3}
		target = []int{7, 3}
		result = addTwoNumbers1(util.IntsToSList(num1), util.IntsToSList(num2))
		resultArr = util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)

		// [2,4,9]
		// [5,6,4,9]
		// [7,0,4,0,1]
		num1 = []int{2, 4, 9}
		num2 = []int{5, 6, 4, 9}
		target = []int{7, 0, 4, 0, 1}
		result = addTwoNumbers1(util.IntsToSList(num1), util.IntsToSList(num2))
		resultArr = util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)

		// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
		// 输出：[8,9,9,9,0,0,0,1]
		num1 = []int{9, 9, 9, 9, 9, 9, 9}
		num2 = []int{9, 9, 9, 9}
		target = []int{8, 9, 9, 9, 0, 0, 0, 1}
		result = addTwoNumbers1(util.IntsToSList(num1), util.IntsToSList(num2))
		resultArr = util.SListToInts(result)
		So(reflect.DeepEqual(resultArr, target), ShouldBeTrue)
	})
}
