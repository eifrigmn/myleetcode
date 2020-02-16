/*
题目：最长公共前缀
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。

	示例 1:
		输入: ["flower","flow","flight"]
		输出: "fl"
	示例 2:
		输入: ["dog","racecar","car"]
		输出: ""
		解释: 输入不存在公共前缀。
	说明:
		所有输入只包含小写字母 a-z 。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
*/
package src

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	var trim = strs[0]
	var s = strs[0]
	var flag = len(s)
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < flag {
			flag = len(strs[i])
		}
		var tr string
		for j := 0; j < flag; j++ {
			if s[j] != strs[i][j] {
				break
			} else {
				tr = tr + string(s[j])
			}
		}
		if len(tr) < len(trim) {
			trim = tr
		}
	}
	return trim
}

// 最优解
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, code := range strs {
		for strings.Index(string(code), prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return string(prefix)
}
