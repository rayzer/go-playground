package main

import "fmt"

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start

/***
Sol Approach v2: Bitmap
Concepts:
    1. Use three array to store the record. (row, col, grid)
    2. For each candidate, we do mask = 1 << int(v), then use row,col,grid to OR that value
    3. Whenever encounter a row,col,grid & mask is true, this mean that we have previously OR this value (duplicate now),
        return false
Time: O(n^2)
Space: O(n)
***/
func isValidSudoku(board [][]byte) bool {
	var row, col, grid [10]int
	n := len(board)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			v := board[r][c]
			if v != '.' {
				mask := 1 << int(v-'0')
				// if has duplicate
				if (row[r]&mask)|(col[c]&mask)|(grid[r/3*3+c/3]&mask) > 0 {
					return false
				}
				row[r] |= mask
				col[c] |= mask
				grid[r/3*3+c/3] |= mask
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
