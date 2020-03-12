package array

import "sort"

func findKthLargest(nums []int, k int) int {
	if !sort.IntsAreSorted(nums) {
		sort.Ints(nums)
	}
	return nums[len(nums)-k]
}

// 使用堆
func findKthLargest2(nums []int, k int) int {
	heap := append([]int{0}, nums[:k]...)
	buildHeap(heap, k)
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[1] {
			// 删除堆顶元素，重新堆化
			heap[1] = nums[i]
			heapify(heap, k, 1)
		}
	}
	return heap[1]

}

func buildHeap(nums []int, n int) {
	for i := n / 2; i > 0; i-- {
		heapify(nums, n, i)
	}
}

func heapify(nums []int, n, i int) {
	for {
		maxPos := i
		if i*2 <= n && nums[i] > nums[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= n && nums[maxPos] > nums[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		nums[i], nums[maxPos] = nums[maxPos], nums[i]
		i = maxPos
	}
}
