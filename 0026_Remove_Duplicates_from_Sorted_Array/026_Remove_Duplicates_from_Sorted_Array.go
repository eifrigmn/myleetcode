package _0026

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var numsLength = 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
		numsLength++
	}
	return numsLength
}

// 参考
func removeDuplicates1(nums []int) int {
	if len(nums) < 2{
		return len(nums)
	}

	int := 0
	for j:= 1; j< len(nums); j ++{
		if nums[int] != nums[j] {
			int++
			nums[int] = nums[j]
		}
	}
	return int +1
}

// 参考
func removeDuplicates2(nums []int) int {
	var count int

	if len(nums) == 0{
		return 0
	}

	count = 1
	for idx:=0 ; idx < len(nums) - 1 ; idx++ {
		if nums[idx] != nums[idx+1]{
			nums[count] = nums[idx+1]
			count += 1
		}
	}
	return count
}