package src

import "testing"

func TestMajorityElement(t *testing.T) {
	var nums = []int{3,2,3}
	t.Log(majorityElement(nums))
	t.Log("*", majorityElement2(nums))
	nums = []int{2, 2, 1, 1, 1, 2, 2}
	t.Log(majorityElement(nums))
	t.Log("*", majorityElement2(nums))
	nums= []int{2, 1, 3, 2, 1, 2, 3, 2, 3}
	t.Log(majorityElement(nums))
	t.Log(majorityElement1(nums))
	t.Log("*", majorityElement2(nums))

}
