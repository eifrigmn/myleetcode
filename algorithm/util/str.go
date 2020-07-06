package util

func StrArry2BytesArry(s []string) []byte {
	var result []byte
	for i := 0; i < len(s); i++ {
		result = append(result, []byte(s[i])...)
	}
	return result
}

func BytesArray2StrArray(b []byte) []string {
	var result []string
	for i := 0; i < len(b); i++ {
		result = append(result, string(b[i]))
	}
	return result
}
