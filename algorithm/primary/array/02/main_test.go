package _02

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	Convey("removeDuplicates", t, func() {
		// [7,1,5,3,6,4]
		prices := []int{7,1,5,3,6,4}
		income := maxProfit(prices)
		So(income, ShouldEqual, 7)
		// [1,2,3,4,5]
		prices = []int{1,2,3,4,5}
		income = maxProfit(prices)
		So(income, ShouldEqual, 4)
		// [7,6,4,3,1]
		prices = []int{7,6,4,3,1}
		income = maxProfit(prices)
		So(income, ShouldEqual, 0)
	})
}
