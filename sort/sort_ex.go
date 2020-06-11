package sort

func BubbleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i:=0;i<len(nums);i++{
		var flag bool
		for j:=0;j<len(nums)-i-1;j++{
			if nums[j+1]<nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 插入排序
func InsertSort(nums []int) {
	//未排序i，已排序j
	for i:=1;i<len(nums);i++{
		var val = nums[i]
		var j = i-1
		for ;j>=0;j--{
			if nums[j]>val {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = val
				break
			}
		}
		nums[j+1] = val
	}
}

