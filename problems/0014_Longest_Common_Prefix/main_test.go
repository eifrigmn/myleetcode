package _0014

import (
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	t.Log(longestCommonPrefix(strs1))
	strs2 := []string{"dog", "racecar", "car"}
	t.Log(longestCommonPrefix(strs2))
	strs3 := []string{"flo", "flowe", "flower", "flowers"}
	t.Log(longestCommonPrefix(strs3))
	strs4 := []string{"aca", "cba"}
	t.Log(longestCommonPrefix(strs4))
	strs5 := []string{"aa", "ab"}
	t.Log(lll(strs5))
}

func lll(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, code := range strs {
		for strings.Index(string(code), prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return string(prefix)
}
