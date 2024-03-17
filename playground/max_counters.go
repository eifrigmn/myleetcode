package main

func MaxCounters(N int, A []int) []int {
	// Implement your solution here
	var maxCounter int
	var result = make([]int, N)
	for i := 0; i < len(A); i++ {
		if A[i] == N+1 {
			fillInResult(result, maxCounter)
			continue
		} else {
			result[A[i]-1]++
			if maxCounter < result[A[i]-1] {
				maxCounter = result[A[i]-1]
			}
		}
	}
	return result
}

func fillInResult(nums []int, val int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = val
	}
}
