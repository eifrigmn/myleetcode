package sort

func insertionSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 1; i < len(nums); i++ {
		var val = nums[i]
		var j = i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = val
				break
			}
		}
		nums[j+1] = val
	}
}
