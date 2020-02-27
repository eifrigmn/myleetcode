package sort

func quickSort(nums []int) []int {
	quickSortC(nums, 0, len(nums)-1)
	return nums
}

func quickSortC(nums []int, lftIdx, rgtIdx int) {
	if lftIdx >= rgtIdx {
		return
	}
	q := partition1(nums, lftIdx, rgtIdx)
	quickSortC(nums, lftIdx, q-1)
	quickSortC(nums, q+1, rgtIdx)
}

// 获取分区点pivot
func partition(nums []int, lftIdx, rgtIdx int) int {
	// 选尾元素为分区点
	pivot := nums[rgtIdx]
	flag := lftIdx
	for i := lftIdx; i < rgtIdx; i++ {
		if nums[i] < pivot {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	nums[flag], nums[rgtIdx] = nums[rgtIdx], nums[flag]
	return flag
}

func partition1(nums []int, lftIdx, rgtIdx int) int {
	// 选头元素为分区点
	pivot := nums[lftIdx]
	flag := lftIdx + 1
	for i := lftIdx + 1; i <= rgtIdx; i++ {
		if nums[i] < pivot {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
	nums[flag-1], nums[lftIdx] = nums[lftIdx], nums[flag-1]
	return flag - 1
}

// 找第K大的元素
func findKthBiggest(nums []int, k int) int {
	kQuickSortC(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func kQuickSortC(nums []int, lftIdx, rgtIdx, k int) {
	if lftIdx >= rgtIdx {
		return
	}

	pivotIdx := kPartition(nums, lftIdx, rgtIdx)
	if pivotIdx == k-1 {
		return
	}
	if pivotIdx < k-1 {
		kQuickSortC(nums, pivotIdx+1, rgtIdx, k)
	} else {
		kQuickSortC(nums, lftIdx, pivotIdx-1, k)
	}
}

func kPartition(nums []int, lftIdx, rgtIdx int) int {
	pivot := nums[rgtIdx]
	flag := lftIdx
	for i := lftIdx; i < rgtIdx; i++ {
		if nums[i] > pivot {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	nums[flag], nums[rgtIdx] = nums[rgtIdx], nums[flag]
	return flag
}
