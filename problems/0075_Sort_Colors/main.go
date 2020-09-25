package _0075

func sortColors(nums []int) {
	var len0, len1 int
	for _, n := range nums {
		if n == 0 {
			len0++
		} else if n == 1 {
			len1++
		}
	}
	idx1 := len0
	idx2 := len0 + len1

	for i := 0; i < len(nums); i++ {
		if i < idx1 {
			nums[i] = 0
		} else if i < idx2 {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

func sortColors1(nums []int) {
	idx0, idx2 := 0, len(nums) - 1
	for i:=0;i <= idx2; {
		if nums[i] == 0 {
			nums[idx0], nums[i] = nums[i], nums[idx0]
			idx0++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[idx2], nums[i] = nums[i], nums[idx2]
			idx2--
		}
	}
}
