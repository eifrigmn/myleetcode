package _0027

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
	}
	return len(nums)
}

// 参考
func removeElement1(nums []int, val int) int {
	p := 0

	for i:= 0; i<len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p += 1
		}
	}

	return p
}
