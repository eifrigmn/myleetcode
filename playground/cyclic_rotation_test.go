package main

import "testing"

func TestCyclicRotation(t *testing.T) {
	nums := []int{1, 1, 2, 3, 5}
	k := -1
	result := CyclicRotation(nums, k)
	t.Log(result)
}
