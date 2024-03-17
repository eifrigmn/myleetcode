package _0001

import (
	"github.com/magiconair/properties/assert"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {
	Convey("twoSum", t, func() {
		//nums := []int{2, 7, 11, 15}
		//target := 22
		//result := _twoSum(nums, target)
		//t.Log(result)
		//assert.Equal(t, result, []int{1, 3})
		nums := []int{3, 3}
		target := 6
		result := _twoSum(nums, target)
		t.Log(result)
		assert.Equal(t, result, []int{0, 1})
	})
}

func TestArray(t *testing.T) {
	arr := [...]int{1, 2, 3}
	t.Log(arr)
}
