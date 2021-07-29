package _02

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("maxProfit", t, func() {
		// [7,1,5,3,6,4]
		prices := []int{7, 1, 5, 3, 6, 4}
		income := maxProfit_dp(prices)
		So(income, ShouldEqual, 7)
		// [1,2,3,4,5]
		prices = []int{1, 2, 3, 4, 5}
		income = maxProfit_greedy(prices)
		So(income, ShouldEqual, 4)
		// [7,6,4,3,1]
		prices = []int{7, 6, 4, 3, 1}
		income = maxProfit_greedy(prices)
		So(income, ShouldEqual, 0)
	})
}
