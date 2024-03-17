## 双指针法：
双指针法可以解决大部分数组问题，如：

[26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)

> 由于原数组已按 **非严格递增** 排列，设置快慢指针，快指针探路，慢指针只记录最终数组中的元素

~~~go
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var slow, fast = 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	nums = nums[:slow+1]
	return slow + 1
}
~~~

[27. 移除元素](https://leetcode.cn/problems/remove-element/)

> 原地移除数组中指定值的元素
>
> 快指针探路，遇到指定值，则跳过，否则，将快指针指向的值赋值给慢指针

~~~go
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != val {
			if fast != slow {
				nums[slow] = nums[fast]
			}
			slow++
		}
		fast++
	}
	nums = nums[:slow]
	return slow
}
~~~

[167. 两数之和 II - 输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/)

> 由于输入的数组有序，因此设置左右指针，分别为数组的头和尾，判断两指针指向的元素的和，小于目标值，则左指针右移，反之则右指针左移

~~~go
func twoSum(numbers []int, target int) []int {
	var lft, rgt = 0, len(numbers) - 1
	for lft < rgt {
		if numbers[lft]+numbers[rgt] > target {
			rgt--
		} else if numbers[lft]+numbers[rgt] < target {
			lft++
		} else {
			return []int{lft + 1, rgt + 1}
		}
	}
	return nil
}
~~~

[344. 反转字符串](https://leetcode.cn/problems/reverse-string/)

> 设置左右指针，分别调换二者指向的字符

~~~go
func reverseString(s []byte) {
	var lft, rgt = 0, len(s) - 1
	for lft < rgt {
		s[lft], s[rgt] = s[rgt], s[lft]
		lft++
		rgt--
	}
}
~~~

[5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)

> 遍历字符串，以遍历点为中心，查找最长的回文字符串
>
> 回文字符串的长度可能为奇数或者偶数，响应的中心点为：
>
> 奇数时：中心为i
>
> 偶数时：中心为i和i+1

~~~go
func longestPalindrome(s string) string {}
~~~



+ [删除有序数组的重复元素](./01/README.md)
+ [移动数组中的0元素](./08/README.md)

## 贪心算法
+ [最佳买卖股票时机](./02/README.md)

## 动态规划法
+ [最佳买卖股票时机](./02/README.md)
