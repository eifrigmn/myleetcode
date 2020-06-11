package _0350

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil 
	}
	nums1 = quickSort(nums1)
	nums2 = quickSort(nums2)
	var p1, p2 int
	var result []int
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}
	return result
}

func quickSort(nums []int) []int {
	quickSortC(nums, 0, len(nums)-1)
	return nums
}

func quickSortC(nums []int, lft, rgt int) {
	if lft >= rgt {
		return
	}
	p := partition(nums, lft, rgt)
	quickSortC(nums, lft, p-1)
	quickSortC(nums, p+1, rgt)
}

func partition(nums []int, lft, rgt int) int {
	flag := lft
	pivot := nums[rgt]
	for i := lft; i < rgt; i++ {
		if nums[i] < pivot {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	nums[flag], nums[rgt] = nums[rgt], nums[flag]
	return flag
}
