package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var block1 = make(map[byte]struct{})
	var block2 = make(map[byte]struct{})
	var block3 = make(map[byte]struct{})
	for i := 0; i < len(board); i++ {
		// 每行数字不重复
		var xTmp = make(map[byte]struct{})
		for j := 0; j < len(board); j++ {
			c := board[i][j]
			if c == byte('.') {
				continue
			}
			if _, ok := xTmp[c]; ok {
				return false
			}
			xTmp[c] = struct{}{}
		}
		// 每列数字不重复
		var yTmp = make(map[byte]struct{})
		for j := 0; j < len(board); j++ {
			c := board[j][i]
			if c == byte('.') {
				continue
			}
			if _, ok := yTmp[c]; ok {
				return false
			}
			yTmp[c] = struct{}{}
		}
		// 分块中的数字不重复
		if i%3 == 0 {
			block1 = make(map[byte]struct{})
			block2 = make(map[byte]struct{})
			block3 = make(map[byte]struct{})
		}
		for j := 0; j < len(board); j++ {
			tmp := block1
			if j >= 3 && j < 6 {
				tmp = block2
			} else if j >= 6 && j < 9 {
				tmp = block3
			}
			c := board[i][j]
			if c == byte('.') {
				continue
			}
			if _, ok := tmp[c]; ok {
				return false
			}
			tmp[c] = struct{}{}
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}

}
