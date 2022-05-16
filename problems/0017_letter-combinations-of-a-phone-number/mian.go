package _017

func letterCombinations(digits string) []string {
	var cache = []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if len(digits) == 0 {
		return []string{}
	}
	var result []string
	str := cache[digits[0]-'0'-1]
	for i:=0;i<len(str);i++{
		result = append(result, string(str[i]))
	}

	for i:=1;i<len(digits);i++{
		var result1 []string
		str := cache[digits[i]-'0'-1]
		for j:=0;j<len(result);j++{
			for k:=0;k<len(str);k++{
				result1 = append(result1, result[j]+string(str[k]))
			}
		}
		result = result1
	}
	return result
}
