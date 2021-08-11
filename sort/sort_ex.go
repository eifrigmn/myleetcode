package sort

func BubbleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		var flag bool
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j+1] < nums[j] {
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
// 数组分为有序区间和无序区间，
// 每次从无序区间拿第一个值，遍历有序区间并插入到有序区间中
func InsertSort(nums []int) {
	//未排序i，已排序j
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

// 归并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// if len(nums) == 2 {
	// 	if nums[0] > nums[1] {
	// 		nums[0], nums[1] = nums[1], nums[0]
	// 	}
	// 	return nums
	// }
	mid := len(nums) / 2
	lft := MergeSort(nums[:mid])
	rgt := MergeSort(nums[mid:])
	return merge1(lft, rgt)
}

func merge1(nums1 []int, nums2 []int) []int {
	var result []int
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			result = append(result, nums1[0])
			nums1 = nums1[1:]
		} else if nums1[0] == nums2[0] {
			result = append(result, nums1[0], nums2[0])
			nums1 = nums1[1:]
			nums2 = nums2[1:]
		} else {
			result = append(result, nums2[0])
			nums2 = nums2[1:]
		}
	}
	if len(nums1) != 0 {
		result = append(result, nums1...)
	}

	if len(nums2) != 0 {
		result = append(result, nums2...)
	}
	return result
}

func Quicksort(nums []int) {
	quickSorts(nums, 0, len(nums)-1)
}

func quickSorts(nums []int, head, tail int) {
	if head >= tail {
		return
	}
	// 指定区间内获取分区点索引
	p := getPivot(nums, head, tail)
	quickSorts(nums, head, p-1)
	quickSorts(nums, p+1, tail)
}

func getPivot(nums []int, head, tail int) int {
	// 以数组头为分区点
	pivot := nums[head]
	pIndex := head + 1
	for i := head + 1; i <= tail; i++ {
		if nums[i] < pivot {
			// 比分区点的值小，则放到分区点左侧
			nums[i], nums[pIndex] = nums[pIndex], nums[i]
			pIndex++
		}
	}
	nums[pIndex-1], nums[head] = nums[head], nums[pIndex-1]
	return pIndex - 1
}
