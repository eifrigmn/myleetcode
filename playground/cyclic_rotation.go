package main

func CyclicRotation(A []int, K int) []int {
	// Implement your solution here
	if len(A) == 0 || K == 0 {
		return A
	}
	K = getK(len(A), K)
	if K == 0 {
		return A
	}
	if K < 0 {
		K = len(A) + K
	}

	return append(A[len(A)-K:], A[:len(A)-K]...)
}

func getK(length, K int) int {
	for K >= length || K <= -length {
		K = K % length
	}
	return K
}
