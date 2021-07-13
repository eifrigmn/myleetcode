package _02

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("卖股票的最佳时机", t, func() {
		// [7,1,5,3,6,4] -> 5
		prices := []int{7, 1, 5, 3, 6, 4}
		So(maxProfit(prices), ShouldEqual, 5)
		// [7,6,4,3,1] -> 0
		prices = []int{7, 6, 4, 3, 1}
		So(maxProfit(prices), ShouldEqual, 0)
	})
}
