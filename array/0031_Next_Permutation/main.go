package array

import (
	"sort"
)

func nextPermutation1(nums []int) {
	preIdx := findPreIndexOfRevspair(nums)
	cIdx := findNextParamIndex(nums, preIdx)
	swapMember(nums, preIdx, cIdx)
	sort.Ints(nums[preIdx+1:])
}

func swapMember(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func findPreIndexOfRevspair(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			return i - 1
		}
	}
	return 0
}

func findNextParamIndex(nums []int, idx int) int {
	for j := len(nums) - 1; j > idx; j-- {
		if nums[j] > nums[idx] {
			return j
		}
	}
	return len(nums) - 1
}

func nextPermutation(nums []int) {
	// 从右向左最长降序数列游标
	idx := longestReverseArrayStart(nums)
	// 将该序列做升序处理
	sort.Ints(nums[idx+1:])
	if idx == -1 {
		return
	}
	// 把序列前的元素，与序列中，第一个大于他的元素互换
	idxToSwap := findForSwap(nums, idx+1)
	// 交换
	nums[idx], nums[idxToSwap] = nums[idxToSwap], nums[idx]
	// tmp := nums[idxToSwap]
	// nums[idxToSwap] = nums[idx-1]
	// nums[idx-1] = tmp
}

// 从右向左寻找最长的降序数列
func longestReverseArrayStart(nums []int) int {
	idx := len(nums) - 2
	for idx >= 0 && nums[idx] >= nums[idx+1] {
		idx--
	}
	return idx
}

// 搜索子序列中第一个大于指定值的元素
// nums 最长降序子序列做升序处理后的序列
// subIdx 最长降序子序列的游标起始值
func findForSwap(nums []int, subIdx int) int {
	target := nums[subIdx-1]
	subIdx--
	r := len(nums) - 1
	for subIdx+1 < r {
		mid := (subIdx + r) / 2
		if target < nums[mid] {
			r = mid
		} else {
			subIdx = mid
		}
	}
	return r
}

// // 参考解
// func nextPermutation1(nums []int) {
// 	if len(nums) <= 1 {
// 		return
// 	}

// 	p := len(nums) - 1

// 	for p > 0 {
// 		if nums[p-1] >= nums[p] {
// 			p--
// 		} else {
// 			break
// 		}
// 	}

// 	// reverse the sub string nums[p:]
// 	l, r := p, len(nums)-1
// 	for l < r {
// 		nums[l], nums[r] = nums[r], nums[l]
// 		l++
// 		r--
// 	}

// 	// find the least greater item than nums[p-1]
// 	if p > 0 {
// 		pivot := nums[p-1]
// 		l, r = p, len(nums)-1
// 		for l < r {
// 			m := (l + r) / 2
// 			if nums[m] <= pivot {
// 				l = m + 1
// 			} else {
// 				r = m
// 			}
// 		}

// 		nums[p-1], nums[l] = nums[l], nums[p-1]
// 	}

// 	return
// }

// // 参考内存占用小
// func nextPermutation2(nums []int) {
// 	if len(nums) <= 1 {
// 		return
// 	}

// 	i := len(nums) - 2
// 	for {
// 		if i < 0 || nums[i] < nums[i+1] {
// 			break
// 		}
// 		i--
// 	}

// 	if i >= 0 {
// 		j := len(nums) - 1
// 		for {
// 			if j == 0 || nums[j] > nums[i] {
// 				break
// 			}
// 			j--
// 		}
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	reverse(nums, i+1)
// }

// func reverse(nums []int, start int) {
// 	i, j := start, len(nums)-1
// 	for {
// 		if i >= j {
// 			return
// 		}
// 		nums[i], nums[j] = nums[j], nums[i]
// 		i++
// 		j--
// 	}
// }

// func isLargest(nums []int) bool {
// 	last := -10000000000
// 	for _, v := range nums {
// 		if v > last {
// 			return false
// 		}
// 		v = last
// 	}
// 	return true
// }

// func quickSort(nums []int) {
// 	if len(nums) <= 1 {
// 		return
// 	}
// 	if len(nums) == 2 {
// 		if nums[0] < nums[1] {
// 			return
// 		}
// 		temp := nums[1]
// 		nums[1] = nums[0]
// 		nums[0] = temp
// 		return
// 	}

// 	pivotIdx := partition(nums)
// 	// recurse
// 	quickSort(nums[:pivotIdx])
// 	quickSort(nums[pivotIdx+1:])

// }

// func partition(a []int) (pivotIdx int) {
// 	pivot := len(a) / 2
// 	left, right := 0, len(a)-1

// 	// Put pivot on the right
// 	a[pivot], a[right] = a[right], a[pivot]

// 	// Pile elements smaller than the pivot on the left
// 	for i := range a {
// 		if a[i] < a[right] {
// 			a[i], a[left] = a[left], a[i]
// 			left++
// 		}
// 	}

// 	// Place the pivot after the last smaller element
// 	a[left], a[right] = a[right], a[left]
// 	return left
// }
