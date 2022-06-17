package _018

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i:=0;i<len(nums)-3;i++{
		target1 := target - nums[i]
		for j:=i+1;j<len(nums)-1;j++{
			var lft, rgt = j+1, len(nums)-1
			for lft < rgt {
				sum := nums[j]+nums[lft]+nums[rgt]
				if sum == target1 {
					var flag bool
					for _, v := range result {
						if v[0] == nums[i] && v[1] == nums[j] && v[2] == nums[lft] && v[3] == nums[rgt] {
							flag = true
						}
					}
					if !flag {
						result = append(result, []int{nums[i], nums[j], nums[lft], nums[rgt]})
					}
					lft++
					rgt--
					for lft < rgt && nums[lft] == nums[lft-1] {
						lft++
					}
					for lft < rgt && nums[rgt] == nums[rgt+1] {
						rgt--
					}
				} else if sum > target1 {
					rgt--
					for lft < rgt && nums[rgt] == nums[rgt+1] {
						rgt--
					}
				} else {
					lft++
					for lft < rgt && nums[lft] == nums[lft-1] {
						lft++
					}
				}

			}
		}
		for ;i<len(nums)-2;i++{
			if nums[i] != nums[i+1] {
				break
			}
		}
	}
	return result
}
