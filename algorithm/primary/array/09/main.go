package _09

func twoSum(nums []int, target int) []int {
	var mp = make(map[int]int, len(nums))
	for i:=0;i<len(nums);i++{
		if j, ok :=mp[target-nums[i]]; ok{
			return []int{j, i}
		}
		mp[nums[i]] = i
	}
	return nil
}
