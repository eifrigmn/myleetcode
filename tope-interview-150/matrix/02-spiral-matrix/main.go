package main

func spiralOrder(matrix [][]int) []int {
	var cnt = 0
	var y = len(matrix)
	var x = len(matrix[0])
	var result = make([]int, x*y)
	for i := 0; i < y; i++ {
		// 上横
		for j := i; j < x-i; j++ {
			if cnt >= x*y {
				break
			}
			result[cnt] = matrix[i][j]
			cnt++
		}
		// 右竖
		for j := i + 1; j < y-i; j++ {
			if cnt >= x*y {
				break
			}
			result[cnt] = matrix[j][x-i-1]
			cnt++
		}

		// 下横
		for j := x - i - 2; j >= i; j-- {
			if cnt >= x*y {
				break
			}
			result[cnt] = matrix[y-i-1][j]
			cnt++
		}

		// 左竖
		for j := y - i - 2; j > i; j-- {
			if cnt >= x*y {
				break
			}
			result[cnt] = matrix[j][i]
			cnt++
		}
	}
	return result
}
