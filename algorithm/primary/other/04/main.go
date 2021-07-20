package _04

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	return generatee(numRows, [][]int{{1}}, []int{1})
}

func generatee(numRows int, pre [][]int, tail []int) [][]int {
	if numRows <= 1 {
		return pre
	}

	return generatee(numRows-1, append(pre, genTail(tail)), genTail(tail))
}

func genTail(num []int) []int {
	if len(num) == 1 {
		return []int{1, 1}
	}
	tail := []int{1}
	for i := 0; i < len(num)-1; i++ {
		tail = append(tail, num[i]+num[i+1])
	}
	return append(tail, 1)
}
