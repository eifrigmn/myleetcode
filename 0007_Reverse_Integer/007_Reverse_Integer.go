package _0007

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	var data int
	flag := 1
	if x < 0 {
		flag = -1
		x = -1 * x
	}
	for x > 0 {
		v := x % 10
		data = data*10 + v
		x /= 10
	}
	if data < math.MinInt32 || data > math.MaxInt32 {
		return 0
	}
	return data * flag
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
