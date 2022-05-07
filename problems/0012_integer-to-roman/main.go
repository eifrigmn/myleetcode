package _012
var flags = []map[int]string{
	{
		10:"X",
		5:"V",
		1: "I",
	},
	{
		10:"C",
		5:"L",
		1: "X",
	},
	{
		10:"M",
		5:"D",
		1:"C",
	},
}

func intToRoman(num int) string {
	var result string
	for i:=0;i<3&&num>0;i++ {
		tail := num%10
		num = num/10
		if tail == 0 {
			continue
		}
		result = getFlag(tail, flags[i])+result
	}
	if num >0 {
		result = repeatChar(num, "M")+result
	}
	return result
}

func getFlag(num int, flag map[int]string) string {
	if num < 4 {
		return repeatChar(num, flag[1])
	} else if num == 4 {
		return flag[1]+flag[5]
	} else if num == 5 {
		return flag[5]
	} else if num < 9 {
		return flag[5]+repeatChar(num-5, flag[1])
	}
	return flag[1] + flag[10]
}

func repeatChar(count int, char string) string {
	var result string
	for i:=0;i<count;i++{
		result += char
	}
	return result
}
