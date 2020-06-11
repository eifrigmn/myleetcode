/*
题目：Fizz Buzz
	写一个程序，输出从 1 到 n 数字的字符串表示。

	1. 如果 n 是3的倍数，输出“Fizz”；

	2. 如果 n 是5的倍数，输出“Buzz”；

	3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

	示例：

		n = 15,

		返回:
		[
    		"1",
    		"2",
    		"Fizz",
    		"4",
    		"Buzz",
    		"Fizz",
    		"7",
    		"8",
    		"Fizz",
    		"Buzz",
    		"11",
    		"Fizz",
    		"13",
    		"14",
    		"FizzBuzz"
		]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fizz-buzz
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package _0412

import "strconv"

func fizzBuzz(n int) []string {
	var result []string
	if n <= 0 {
		return nil
	}
	for i:=1;i<=n;i++{
		if i % 15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i % 3 == 0 {
			result = append(result, "Fizz")
		} else if i % 5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// 参考
func fizzBuzz1(n int) []string {
	result := make([]string, 0, n)
	var i int
	for i = 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}
		if i % 3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		if i % 5 == 0 {
			result = append(result, "Buzz")
			continue
		}
		result = append(result, strconv.Itoa(i))
	}
	return result
}

func fizzBuzz2(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		newValue := ""

		if i%3 == 0 {
			newValue += "Fizz"
		}
		if i%5 == 0 {
			newValue += "Buzz"
		}

		if newValue == "" {
			newValue = strconv.Itoa(i)
		}

		result = append(result, newValue)
	}

	return result
}