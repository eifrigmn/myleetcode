package _012

func intToRoman1(num int) string {
	l1 := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	l2 := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	res := ""

	for i, v := range l2 {
		for num >= v {
			res += l1[i]
			num -= v
		}
	}

	return res
}
