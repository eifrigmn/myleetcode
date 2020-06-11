package _0056

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort(intervals)
	var result [][]int
	pivot := intervals[0]
	for i := 1; i < len(intervals); i++ {
		flag := intervals[i]
		if flag[0] <= pivot[1] {
			if pivot[1] < flag[1] {
				pivot[1] = flag[1]
			}
		} else {
			result = append(result, pivot)
			pivot = intervals[i]
		}
	}
	result = append(result, pivot)
	return result
}

func sort(nums [][]int) [][]int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums [][]int, lft, rgt int) {
	if lft >= rgt {
		return
	}
	p := partition(nums, lft, rgt)
	quickSort(nums, lft, p-1)
	quickSort(nums, p+1, rgt)
}

func partition(nums [][]int, lft, rgt int) int {
	flag := lft
	pivot := nums[rgt]
	for i := lft; i < rgt; i++ {
		if nums[i][0] < pivot[0] || (nums[i][0] == pivot[0] && nums[i][1] < pivot[1]) {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	nums[flag], nums[rgt] = nums[rgt], nums[flag]
	return flag
}
