package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}

	var solve func([][]byte) bool
	solve = func([][]byte) bool {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == '.' {
					for c := byte('1'); c <= byte('9'); c++ {
						if isValid(board, i, j, c) {
							board[i][j] = c

							if solve(board) {
								return true
							} else {
								board[i][j] = '.'
							}
						}
					}
					return false
				}
			}
		}
		return true
	}
	solve(board)
	fmt.Println(board)
}

func isValid(board [][]byte, i, j int, ci byte) bool {
	//for i := 0; i < len(board); i++ {
	//	if board[i][col] != '.' && board[i][col] == c {
	//		return false
	//	}
	//	if board[row][i] != '.' && board[row][i] == c {
	//		return false
	//	}
	//	if board[3*(row/3)+i/3][3*(col/3)+i%3] != '.' && board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
	//		return false
	//	}
	//}
	var row, col, grid [10]int
	n := len(board)
	board[i][j] = ci

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			v := board[r][c]
			if v != '.' {
				mask := 1 << int(v-'0')
				// if has duplicate
				if (row[r]&mask)|(col[c]&mask)|(grid[r/3*3+c/3]&mask) > 0 {
					board[i][j] = '.'
					return false
				}
				row[r] |= mask
				col[c] |= mask
				grid[r/3*3+c/3] |= mask
			}
		}
	}
	board[i][j] = '.'
	return true
}

// @lc code=end

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
}
