package _04

func rob(nums []int) int {
	var max0, max1 int
	max1 = nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := max0
		if max1 > max0 {
			tmp = max1
		}
		max1 = max0 + nums[i]
		max0 = tmp
	}

	if max0 > max1 {
		return max0
	}

	return max1

}
