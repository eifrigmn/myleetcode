package _011

func maxArea(height []int) int {
	var l = 0
	var r = len(height) - 1
	var result int
	for l < r {
		result = getMax(result, (r-l)*getMin(height[l], height[r]))
		if height[l] <= height[r]{
			l++
		} else {
			r--
		}
	}
	return result
}

func getMin(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
