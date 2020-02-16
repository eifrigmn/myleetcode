package array

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {
	Convey("twoSum", t, func() {
		nums := []int{2, 7, 11, 15}
		target := 22
		fmt.Println(twoSum(nums, target))
	})
}

func TestString(t *testing.T) {
	s := "IVXLCDM"
	for _, c := range s {
		t.Log(fmt.Sprintf("%c", c), c)
	}
}
