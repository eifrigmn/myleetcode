/*
题目：实现 strStr()

	实现 strStr() 函数。
	给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

	示例 1:
	输入: haystack = "hello", needle = "ll"
	输出: 2
	示例 2:
	输入: haystack = "aaaaa", needle = "bba"
	输出: -1
	说明:
		当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
		对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
*/
package src

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) <len(needle) {
		return -1
	}
	for i:=0;i<len(haystack);i++ {
		if len(haystack[i:]) <len(needle) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// 参考
func strStr1(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)
	if lh < ln {
		return -1
	}
	if ln == 0 {
		return 0
	}

	flag := 0
	for flag <= lh-ln {
		if haystack[flag:flag+ln] == needle {
			return flag
		} else {
			flag += ln
			if flag >= lh {
				return -1
			}
			has := false
			for i := ln - 1; i >= 0; i-- {
				if needle[i:i+1] == haystack[flag:flag+1] {
					flag -= i
					has = true
					break
				}
			}
			if !has {
				flag++
			}
		}
	}
	return -1
}