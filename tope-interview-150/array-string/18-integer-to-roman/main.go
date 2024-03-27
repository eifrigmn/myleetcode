package main

func intToRoman(num int) string {
	var result string
	for num > 0 {
		item, cnt := getMaxItem(num)
		num = num - cnt
		result = result + item
	}
	return result
}

func getMaxItem(num int) (string, int) {
	if num >= 1000 {
		return "M", 1000
	} else if num >= 900 {
		return "CM", 900
	} else if num >= 500 {
		return "D", 500
	} else if num >= 400 {
		return "CD", 400
	} else if num >= 100 {
		return "C", 100
	} else if num >= 90 {
		return "XC", 90
	} else if num >= 50 {
		return "L", 50
	} else if num >= 40 {
		return "XL", 40
	} else if num >= 10 {
		return "X", 10
	} else if num >= 9 {
		return "IX", 9
	} else if num >= 5 {
		return "V", 5
	} else if num >= 4 {
		return "IV", 4
	} else if num >= 1 {
		return "I", 1
	}
	return "", 0
}
