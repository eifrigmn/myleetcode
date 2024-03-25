package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("O(1) 时间插入、删除和获取随机元素", t, func() {
		this := Constructor()

		resp := this.Insert(1)
		So(resp, ShouldBeTrue)

		resp = this.Remove(2)
		So(resp, ShouldBeFalse)

		resp = this.Insert(2)
		So(resp, ShouldBeTrue)

		data := this.GetRandom()
		So(data, ShouldBeIn, []int{1, 2})

		resp = this.Remove(1)
		So(resp, ShouldBeTrue)

		resp = this.Insert(2)
		So(resp, ShouldBeFalse)

		data = this.GetRandom()
		So(data, ShouldEqual, 2)
	})
}
