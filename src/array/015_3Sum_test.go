package array

import "testing"

func TestThresSum1(t *testing.T) {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(nums))
	nums = []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	t.Log(threeSum(nums))
}
