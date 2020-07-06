package util

import (
	"fmt"
	"testing"
)

func TestStrArry2BytesArry(t *testing.T) {
	var s = []string{"h", "e", "l", "l", "o"}
	b := StrArry2BytesArry(s)
	fmt.Println(string(b))
}
