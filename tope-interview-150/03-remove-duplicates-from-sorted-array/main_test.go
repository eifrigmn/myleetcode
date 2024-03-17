package main

import (
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var nums = []int{1, 1, 2}
	res := removeDuplicates(nums)
	assert.Equal(t, res, 2)
	assert.Equal(t, reflect.DeepEqual(nums[:res], []int{1, 2}), true)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res = removeDuplicates(nums)
	assert.Equal(t, res, 5)
	assert.Equal(t, reflect.DeepEqual(nums[:res], []int{0, 1, 2, 3, 4}), true)
}
