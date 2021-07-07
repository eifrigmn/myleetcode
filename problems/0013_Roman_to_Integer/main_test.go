package _0013

import (
	"strings"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	var s1 = "III"
	t.Log(romanToInt(s1))
	var s2 = "IV"
	t.Log(romanToInt(s2))
	var s3 = "IX"
	t.Log(romanToInt(s3))
	var s4 = "LVIII"
	t.Log(romanToInt(s4))
	var s5 = "MCMXCIV"
	t.Log(romanToInt(s5))
}

func TestStr(t *testing.T) {
	s1 := "flower"
	s2 := "ower"
	t.Log(strings.Trim(s1,s2))
}
