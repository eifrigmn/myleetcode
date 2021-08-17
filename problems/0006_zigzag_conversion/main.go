package _0006

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) == 1 {
		return s
	}
	ss := convertToSlice(s, numRows)
	var result string
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(ss); j++ {
			if len(ss[j]) <= i {
				break
			}
			if ss[j][i] == ' ' {
				continue
			}
			result = fmt.Sprintf("%s%c", result, ss[j][i])
		}
	}
	return result
}

func convertToSlice(s string, numRows int) []string {
	var ss []string
	var flag int
	for len(s) > 0 {
		if flag == 0 || flag == numRows-1 {
			if len(s) > numRows {
				ss = append(ss, s[:numRows])
				s = s[numRows:]
			} else {
				ss = append(ss, s)
				break
			}
		} else {
			var tmp string
			for i := numRows - 1; i >= 0; i-- {
				if i != flag {
					tmp = tmp + " "
				} else {
					tmp = tmp + string(s[0])
				}
			}
			ss = append(ss, tmp)
			if len(s) == 1 {
				break
			} else {
				s = s[1:]
			}
		}
		if flag == numRows-1 {
			flag = 1
		} else {
			flag++
		}
	}
	return ss
}

// 示例
// 数组不存储中间状态的切割后的数组
// 长度为numRows的切片存储转换后的子字符串
// 最后字符串合并
func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	dir, turn := -1, 0
	for _, v := range s {
		rows[turn] += string(v)
		if (turn+1)%numRows == 0 || turn == 0 {
			dir = -dir
		}
		turn += dir
	}
	return strings.Join(rows, "")
}
