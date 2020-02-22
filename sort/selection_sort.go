package sort 

func selectionSort(nums []int) {
	for i := range nums {
		var minIdx = i 
		for j := i+1; j <len(nums); j++{
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}