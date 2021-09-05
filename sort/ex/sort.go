package sort_ex

// 冒泡排序
func bubble(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 选择排序
// 不断从未排序区间选择最小的元素放置到已排序区间的末尾
func selection(nums []int) {
	for sortedTailIdx := 0; sortedTailIdx < len(nums); sortedTailIdx++ {
		var minIdx = sortedTailIdx
		for i := sortedTailIdx + 1; i < len(nums); i++ {
			if nums[minIdx] > nums[i] {
				minIdx = i
			}
		}
		nums[sortedTailIdx], nums[minIdx] = nums[minIdx], nums[sortedTailIdx]
	}
}

// 插入排序
// 每次从未排序区间取一个数插入到已排序区间中
func insert(nums []int) {
	if len(nums) < 2 {
		return
	}
	// i表示未排序区间的第一个值
	// 0~j为有序区间
	// 因此，相当于元素nums[i]在区间0~j范围内向上冒泡
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 归并排序
// 将一个数组等分，分别对分割后的两个数组排序，然后将其合并
// 如此递归，数组被递归等分，最终为比较两个最小长度数组的值并按顺序合并
func merge(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	lft := nums[:mid]
	rgt := nums[mid:]
	return m(merge(lft), merge(rgt))
}

func m(nums1, nums2 []int) []int {
	var result []int
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] > nums2[0] {
			result = append(result, nums2[0])
			nums2 = nums2[1:]
		} else {
			result = append(result, nums1[0])
			nums1 = nums1[1:]
		}
	}
	if len(nums1) > 0 {
		result = append(result, nums1...)
	}

	if len(nums2) > 0 {
		result = append(result, nums2...)
	}
	return result
}

// 快速排序
// 在数组中找一个分区点p
// 遍历数组，大于p的放到p的右边，小于p的放到p的左边
// 如此无限分割
// 直到子序列的长度为1
func quickSort(nums []int) {
	quickSoctC(nums, 0, len(nums)-1)
}

func quickSoctC(nums []int, head, tail int) {
	if head >= tail {
		return
	}
	p := getPivort(nums, head, tail)
	quickSoctC(nums, head, p-1)
	quickSoctC(nums, p+1, tail)
}

func getPivort(nums []int, head, tail int) int {
	var pivort = nums[head]
	var nextPIdx = head + 1
	for i := head + 1; i <= tail; i++ {
		if nums[i] < pivort {
			nums[i], nums[nextPIdx] = nums[nextPIdx], nums[i]
			nextPIdx++
		}
	}
	nums[nextPIdx-1], nums[head] = nums[head], nums[nextPIdx-1]
	return nextPIdx - 1
}

// 希尔排序
// 堆排序
// 计数排序
// 桶排序
// 基数排序
