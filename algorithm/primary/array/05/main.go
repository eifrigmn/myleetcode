package _05

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var mp = make(map[int]int, len(nums))
	var result int
	for i:=0;i<len(nums);i++{
		if _, ok := mp[nums[i]];ok {
			result -= nums[i]
		} else {
			result += nums[i]
			mp[nums[i]] = 1
		}
	}
	return result
}
