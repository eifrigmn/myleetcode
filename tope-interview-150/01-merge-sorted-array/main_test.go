package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{1, 2, 2, 3, 5, 6})

	nums1 = []int{1}
	nums2 = []int{}
	m = 1
	n = 0
	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{1})

	nums1 = []int{0}
	nums2 = []int{1}
	m = 0
	n = 1
	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{1})

	nums1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	nums2 = []int{1, 2, 2}
	m = 6
	n = 3
	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{-1, 0, 0, 1, 2, 2, 3, 3, 3})

	nums1 = []int{-1, -1, 0, 0, 0, 0}
	nums2 = []int{-1, 0}
	m = 4
	n = 2
	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{-1, -1, -1, 0, 0, 0})
}
