package _0015

import "sort"

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		// 避免添加重复结果
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}
		lftIndex := i + 1
		rgtIndex := len(nums) - 1
		for lftIndex < rgtIndex {
			sum := nums[i] + nums[lftIndex] + nums[rgtIndex]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[lftIndex], nums[rgtIndex]})
			loop:
				for lftIndex < rgtIndex {
					switch {
					case nums[lftIndex] == nums[lftIndex+1]:
						lftIndex++
					case nums[rgtIndex] == nums[rgtIndex-1]:
						rgtIndex--
					default:
						lftIndex++
						rgtIndex--
						break loop
					}
				}
			} else if sum > 0 {
				rgtIndex--
			} else {
				lftIndex++
			}
		}
	}
	return res
}

// 参考
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func threeSum1(nums []int) [][]int {
	res := [][]int{}
	sort.Sort(IntSlice(nums))

	l := len(nums)
	var sum int
	for i := 0; i < l; {
		left := i + 1
		right := l - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
			}

			if sum > 0 {
				right--
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else {
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			}
		}

		i++
		for i < l && nums[i] == nums[i-1] {
			i++
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		if nums[i] > 0 {
			break
		}
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// 较小的 l 需要变大
				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
