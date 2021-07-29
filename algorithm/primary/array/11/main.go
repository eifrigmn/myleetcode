package _11

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
		}
	}

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			matrix[i][j], matrix[l-1-j][l-1-i] = matrix[l-1-j][l-1-i], matrix[i][j]
		}
	}
}

func rotate1(matrix [][]int) {
	// 对角线互换
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 左右对称互换
	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
		}
	}
}
