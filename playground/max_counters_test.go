package main

import "testing"

func TestMaxCounters(t *testing.T) {
	nums := []int{3, 4, 4, 6, 1, 4, 4}
	N := 5
	result := MaxCounters(N, nums)
	t.Log(result)

	nums = []int{2, 1, 1, 2, 1}
	N = 1
	result = MaxCounters(N, nums)
	t.Log(result)

	nums = []int{1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 11}
	N = 10
	result = MaxCounters(N, nums)
	t.Log(result)
}
