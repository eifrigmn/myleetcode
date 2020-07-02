package _10

func isValidSudoku(board [][]byte) bool {
	var (
		mph = make(map[int]map[string]int, 9)
		mp3 = make(map[int]map[string]int, 9)
	)
	for i:=0;i<len(board);i++{
		var mp = make(map[string]int, 9)

		for j :=0;j<len(board[i]);j++{
			if string(board[i][j]) == "." {
				continue
			}
			if _, ok := mp[string(board[i][j])]; ok {
				return false
			}
			mp[string(board[i][j])] = 1
			// 列
			if mph[j] == nil {
				mph[j] = make(map[string]int, 9)
			}

			if _, ok := mph[j][string(board[i][j])]; ok {
				return false
			}

			mph[j][string(board[i][j])] = 1
			// 3x3
			idx := (i/3)*3 + j/3
			if mp3[idx] == nil {
				mp3[idx] = make(map[string]int, 9)
			}

			if _, ok := mp3[idx][string(board[i][j])]; ok {
				 return false
			}
			mp3[idx][string(board[i][j])] = 1
		}

	}
	return true
}

func isValidSudoku1(board [][]byte) bool {
	for i := range board {
		x := make(map[byte]bool)
		for j := range board {
			if board[i][j] != []byte(".")[0] {
				if x[board[i][j]] {
					return false
				}
				x[board[i][j]] = true
			}
		}
	}

	for i := 0; i < len(board); i++ {
		y := make(map[byte]bool)
		for j := 0; j < len(board); j++ {
			if board[j][i] != []byte(".")[0] {
				if y[board[j][i]] {
					return false
				}
				y[board[j][i]] = true
			}
		}
	}

	for i := 0; i < len(board); {
		for j := 0; j < len(board); {
			if !check(i, j, board) {
				return false
			}
			j += 3
		}
		i += 3
	}

	return true
}

func check(x, y int, board [][]byte) bool {
	z := make(map[byte]bool)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if board[i][j] != []byte(".")[0] {
				if z[board[i][j]] {
					return false
				}
				z[board[i][j]] = true
			}
		}
	}
	return true
}