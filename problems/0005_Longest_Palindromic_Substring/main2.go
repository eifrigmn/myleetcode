package _0005

//func _longestPalindrome(s string) string {
//	if len(s) <= 1 {
//		return s
//	}
//	// 找中点
//	mid := len(s) / 2
//	// 回文字符串，应该首字母和尾字母一致
//	var head int
//	var str string
//	for head <= mid {
//		if len(s)-head < len(str) {
//			return str
//		}
//		h := head
//		t := len(s) - 1
//		tt := t
//		for t > h {
//			if s[h] == s[t] {
//				h++
//				t--
//				if t-h < 2 {
//					break
//				}
//			} else {
//				h = head
//				t--
//				tt = t
//				continue
//			}
//		}
//		if s[h] == s[t] && t-h < 2 {
//			if tt-head+1 > len(str) {
//				str = s[head : tt+1]
//			}
//		}
//		head++
//	}
//	return str
//}
