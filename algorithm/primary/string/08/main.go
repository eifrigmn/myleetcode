package _08

import "strconv"

func countAndSay(n int) string {
	var result string
	if n == 1 {
		return "1"
	} else {
		prevResult := countAndSay(n-1)
		flag := prevResult[0]
		if len(prevResult) == 1 {
			return "1"+prevResult
		}

		var count = 1

		for i:=1;i<len(prevResult);i++{
			if prevResult[i] == flag {
				count++
				continue
			}

			countStr := strconv.Itoa(count)
			result = result + countStr + string(flag)
			flag = prevResult[i]
			count = 1
		}

		countStr := strconv.Itoa(count)
		result = result + countStr + string(flag)
	}
	return result
}

