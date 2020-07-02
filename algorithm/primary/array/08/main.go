package _08

func moveZeroes(nums []int)  {
	if len(nums) == 1 {
		return
	}
	var zidx int
	for i:=0;i<len(nums);i++{
		if nums[i] != 0 {
			nums[i], nums[zidx] = nums[zidx], nums[i]
			zidx++
		}
	}
}
