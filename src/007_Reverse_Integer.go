/*
题目：整数反转
	给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

	示例 1:
		输入: 123
		输出: 321
 	示例 2:
		输入: -123
		输出: -321
	示例 3:
		输入: 120
		输出: 21
	注意:
		假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
*/
package src

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	var data int
	for true {
		v := x % 10
		x /= 10
		if v == 0 && x == 0 {
			break
		}
		data = data*10 + v

	}
	if data < math.MinInt32 || data > math.MaxInt32 {
		return 0
	}
	return data
}

// 最优解
func reverse1(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}
	item := strconv.Itoa(x)
	ret := make([]rune, len(item))
	index := len(item) - 1

	for _, v := range item {
		ret[index] = v
		index--
	}

	num, error := strconv.Atoi(string(ret))
	if error != nil {
		return 0
	}
	num = flag * num

	intMax := int(^uint32(0) >> 1)
	intMin := ^intMax

	if intMin < num && num < intMax {
		return num
	}

	return 0
}
