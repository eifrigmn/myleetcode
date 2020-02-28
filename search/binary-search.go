package search

// 查找第一个大于等于给定值的元素
// nums: 原始数据
// val: 给定值
// 返回查找到的元素索引，未找到则返回-1
func bSearchUp(nums []int, val int) int {
	var lft, rgt = 0, len(nums) - 1
	for lft <= rgt {
		mid := lft + ((rgt - lft) >> 1)
		if nums[mid] < val {
			lft = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < val {
				return mid
			} else {
				rgt = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个等于值的元素
func bSearchTail(nums []int, val int) int {
	var lft, rgt = 0, len(nums) - 1
	for lft <= rgt {
		mid := lft + ((rgt - lft) >> 1)
		if nums[mid] < val {
			lft = mid + 1
		} else if nums[mid] > val {
			rgt = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != val {
				return mid
			} else {
				lft = mid + 1
			}
		}
	}
	return -1
}

// 第一个等于给定值的元素
func bSearchHead(nums []int, val int) int {
	var lft, rgt = 0, len(nums) - 1
	for lft <= rgt {
		mid := lft + ((rgt - lft) >> 1)
		if nums[mid] < val {
			lft = mid + 1
		} else if nums[mid] > val {
			rgt = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != val {
				return mid
			} else {
				rgt = mid - 1
			}
		}
	}
	return -1
}

// 最后一个小于等于给定值的元素
func bSearchDown(nums []int, val int) int {
	var lft, rgt = 0, len(nums) - 1
	for lft <= rgt {
		mid := lft + ((rgt - lft) >> 1)
		if nums[mid] > val {
			rgt = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > val {
				return mid
			} else {
				lft = mid + 1
			}
		}
	}
	return -1
}
