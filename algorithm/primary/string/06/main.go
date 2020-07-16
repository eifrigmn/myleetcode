package _06

import "math"

func myAtoi(str string) int {
	var result, i int
	var isNegative bool
	for len(str) > 0 {
		if str[0] == ' ' {
			if len(str) == 1 {
				return 0
			} else {
				str = str[1:]
				continue
			}
		} else if str[0] == '-' || str[0] == '+'{
			if len(str) == 1 {
				return 0
			} else if str[1] < '0' || str[1] > '9'{
				return 0
			} else {
				if str[0] == '-' {
					isNegative = true
				}
				str = str[1:]
				continue
			}
		} else if str[0] >= '0' && str[0] <= '9' {
			if i >= len(str) || str[i] < '0' || str[i] > '9' {
				break
			} else {
				if isNegative && -result < math.MinInt32{
					return math.MinInt32
				} else if !isNegative && result > math.MaxInt32{
					return math.MaxInt32
				}
				num := str[i] - '0'
				result = result*10+int(num)
				i++
				continue
			}
		} else {
			return 0
		}
	}

	if isNegative {
		result = -result
	}

	if result < math.MinInt32 {
		return math.MinInt32
	} else if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return result
}
