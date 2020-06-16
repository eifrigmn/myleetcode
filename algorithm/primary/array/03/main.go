package _03

func rotate(nums []int, k int)  {
	if len(nums) < 2{
		return
	}
	for i:=0;i<k;i++{
		tmp := nums[len(nums)-1]
		for i:=len(nums)-2;i>=0;i--{
			nums[i+1] = nums[i]
		}
		nums[0] = tmp
	}
}

func rotate1(nums []int, k int)  {
	if len(nums)<=1{
		return
	}
	//nums[0] = 100
	k = k%len(nums)
	result := append(nums[len(nums)-k:],nums[:len(nums)-k]...)
	for idx,item :=range result{
		nums[idx] = item
	}
}
