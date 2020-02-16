/*
题目：下一个排列
	实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
	如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
	必须原地修改，只允许使用额外常数空间。

	以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
		1,2,3 → 1,3,2
		3,2,1 → 1,2,3
		1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package array

import (
	"sort"
)

/*
说明：
	字典序
		给定多个字符，可以按照任意顺序进行排列，所有排列称为全排列。

		每一种排列对应一个字符串，如果这些字符串按照字符串大小的顺序进行排序，那么就这种排序是基于字典序的全排列。

		比如给定1，2，3，则他们基于字典序的全排列为：
		123> 132 > 213 >231 > 312 > 321
*/

// 结合一个现实的应用：
// 	可以将已有排列组合成一串数字，下一个排列的数字要比上一个排列的数字大
//
// 找规律，以序列`1, 5, 8, 4, 7, 6, 5, 3, 1`为例：
// 	如果存在下一个排列，则原始排列中必然存在升序排列，才能满足下一个排列中必然出现新的降序排列
// 	因此可以认为我们需要在原始排列中寻找排列顺序变更的点，即原始序列由升序转向降序的数字
// 	由于需要寻找的是下一个排列，即上一个排列合成的数字要比下一个排列合成的数字小，因此我们选定从右向左寻找变更点
//
// 算法步骤：
//	1. 从右向左寻找第一组长度为2的降序数组:
//		1, 5, 8, [4, 7], 6, 5, 3, 1
//  2. 选定需要被提换的数字为4，从右向左寻找第一个比4大的数字，即5
//                        !
//		1, 5, 8, 4, 7, 6, 5, 3, 1
//               ^
//  3. 将5与4调换位置
//                        !
//		1, 5, 8, 5, 7, 6, 4, 3, 1
// 				 ^
// 	4. 换位后，将`5`后面的数组做升序处理
// 		1, 5, 8, 5, 1, 3, 4, 6, 7
//               ^
func nextPermutation(nums []int) {
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
	return len(nums)-1
}

// 参考解
func nextPermutation1(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	p := len(nums) - 1

	for p > 0 {
		if nums[p-1] >= nums[p] {
			p--
		}else{
			break
		}
	}

	// reverse the sub string nums[p:]
	l, r := p, len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	// find the least greater item than nums[p-1]
	if p > 0 {
		pivot := nums[p-1]
		l, r = p, len(nums)-1
		for l < r {
			m := (l+r)/2
			if nums[m] <= pivot{
				l = m+1
			}else{
				r = m
			}
		}

		nums[p-1], nums[l] = nums[l], nums[p-1]
	}

	return
}

// 参考内存占用小
func nextPermutation2(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 2;
	for {
		if  i < 0 || nums[i] < nums[i+1]  {
			break
		}
		i--
	}

	if (i >= 0) {
		j := len(nums) - 1
		for {
			if j == 0 || nums[j] > nums[i] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	i, j := start, len(nums) -1
	for {
		if i >= j {
			return
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func isLargest(nums []int) bool {
	last := -10000000000
	for _, v := range nums {
		if v > last {
			return false
		}
		v = last
	}
	return true
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	if len(nums) == 2{
		if nums[0] < nums[1] {
			return
		}
		temp := nums[1]
		nums[1] = nums[0]
		nums[0] = temp
		return
	}

	pivotIdx := partition(nums)
	// recurse
	quickSort(nums[:pivotIdx])
	quickSort(nums[pivotIdx + 1:])

}

func partition(a []int) (pivotIdx int) {
	pivot := len(a) / 2
	left, right := 0, len(a) - 1

	// Put pivot on the right
	a[pivot], a[right] = a[right], a[pivot]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]
	return left
}
