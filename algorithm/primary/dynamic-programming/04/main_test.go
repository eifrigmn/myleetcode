package _04

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRob(t *testing.T) {
	Convey("打家劫舍", t, func() {
		// [1,2,3,1] -> 4
		//	i	nums[i]		max0		max1
		// 	0	1			0			1=nums[0]
		// 	1	2			1=max(0,1)	2=0+2
		//	2	3			2=max(1,2)	4=1+3
		//	3	1			4=max(2,4)	3=2+1
		nums := []int{1, 2, 3, 1}
		So(rob(nums), ShouldEqual, 4)
		// [2,7,9,3,1] -> 12
		nums = []int{2, 7, 9, 3, 1}
		So(rob(nums), ShouldEqual, 12)
		// [7,2,3,9,1] -> 16
		//	i	nums[i]		max0			max1
		// 	0	7			0				7=nums[0]
		// 	1	2			7=max(0,7)		2=0+2
		//	2	3			7=max(7,2)		10=7+3
		//	3	9			10=max(7,10)	16=7+9
		//	4	1			16=max(16,10)	11=10+1
		nums = []int{7, 2, 3, 9, 1}
		So(rob(nums), ShouldEqual, 16)
	})
}
