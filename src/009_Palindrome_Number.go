/*
题目：回文数
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

	示例 1:
		输入: 121
		输出: true
	示例 2:
		输入: -121
		输出: false
		解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
	示例 3:
		输入: 10
		输出: false
		解释: 从右向左读, 为 01 。因此它不是一个回文数。
	进阶:
		你能不将整数转为字符串来解决这个问题吗？
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
*/
package src

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revertedNumber int
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10
}

// 最优解
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	} else {
		r := 0
		t := x
		for x > 0 {
			r = r*10 + x%10
			x /= 10
		}
		if r == t {
			return true
		} else {
			return false
		}
	}
}
