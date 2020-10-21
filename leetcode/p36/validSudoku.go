package main

import "fmt"

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	set := make(map[string]struct{})
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			number := board[i][j]
			if number != '.' {
				if _, exist := set[string(number)+" in row "+string(i)]; exist {
					return false
				} else {
					set[string(number)+" in row "+string(i)] = struct{}{}
				}
				if _, exist := set[string(number)+" in column "+string(j)]; exist {
					return false
				} else {
					set[string(number)+" in column "+string(j)] = struct{}{}
				}
				if _, exist := set[string(number)+" in block "+string(i/3)+"-"+string(j/3)]; exist {
					return false
				} else {
					set[string(number)+" in block "+string(i/3)+"-"+string(j/3)] = struct{}{}
				}
			}
		}
	}
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
	fmt.Print(isValidSudoku(board))
}
