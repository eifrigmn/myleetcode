package sort

func merge(A []int, m int, B []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		A = append(A[:m], B[:n]...)
		return 
	}

	if A[m-1] <= B[0] {
		A = append(A[:m], B[:n]...)
		return
	}

	if B[n-1] <= A[0] {
		copy(A[m:], A[:m])
		copy(A[:m], B[:n])
		return
	}
	A = append(A[:m], B[:n]...)
	quick(A, 0, m+n-1)
	// rgt := m
	// for len(B) > 0 {
	// 	pos := getInstPos(A, rgt, B[len(B)-1])
	// 	var temp = []int{B[len(B)-1]}
	// 	temp = append(temp, A[pos:m]...)
	// 	A = append(A[:pos], temp...)
	// 	B = B[:len(B)-1]
	// 	m++
	// 	rgt = pos
	// }
}

// func getInstPos(nums []int, rgt, val int) int {
// 	var i = 0
// 	for i < rgt && nums[i] < val {
// 		i++
// 	}
// 	return i
// }

func quick(nums []int, lft, rgt int) {
	if lft >= rgt {
		return
	}
	p := partition(nums, lft, rgt)
	quick(nums, lft, p-1)
	quick(nums, p+1, rgt)
}

func partition(nums []int, lft, rgt int) int {
	pivot := nums[rgt]
	flag := lft
	for i := lft; i < rgt; i++ {
		if nums[i] < pivot {
			nums[flag], nums[i] = nums[i], nums[flag]
			flag++
		}
	}
	nums[flag], nums[rgt] = nums[rgt], nums[flag]
	return flag
}