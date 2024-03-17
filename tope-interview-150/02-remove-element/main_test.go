package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	size := removeElement1(nums, val)
	assert.Equal(t, nums, []int{2, 2})
	assert.Equal(t, size, 2)

	nums = []int{3, 3}
	val = 3
	size = removeElement1(nums, val)
	assert.Equal(t, nums, []int{})
	assert.Equal(t, size, 0)

	nums = []int{3}
	val = 3
	size = removeElement1(nums, val)
	assert.Equal(t, nums, []int{})
	assert.Equal(t, size, 0)
}
