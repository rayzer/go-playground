package main

import (
	"fmt"
	"reflect"
	"strings"
)

func solveNQueens(n int) [][]string {
	cols := make(map[int]bool)
	pie := make(map[int]bool)
	na := make(map[int]bool)
	var results [][]string

	if n < 1 {
		return results
	}
	var DFS func(n, row int, cur_stnate []int)

	DFS = func(n, row int, cur_state []int) {
		if row >= n {
			results = append(results, generateBoard(cur_state))
		}

		for col := 0; col < n; col++ {
			if true == cols[col] || true == pie[row+col] || true == na[row-col] {
				continue //not able to place
			}

			if reflect.DeepEqual([]int{5, 0, 2, 4, 6, 3}, cur_state) {
				fmt.Println(col, row, cols, pie, na)
			}
			//place it
			cols[col], pie[row+col], na[row-col] = true, true, true
			new_cur_state := append(cur_state, col)
			DFS(n, row+1, new_cur_state)

			//reverse
			new_cur_state = new_cur_state[:len(new_cur_state)-1]
			cols[col], pie[row+col], na[row-col] = false, false, false
		}
	}

	DFS(n, 0, []int{})
	return results
}

func generateBoard(result []int) (board []string) {
	for i := 0; i < len(result); i++ {
		row := strings.Repeat(".", result[i]) + "Q" + strings.Repeat(".", len(result)-result[i]-1)
		board = append(board, row)
	}
	return
}

func main() {
	fmt.Println(solveNQueens(7))
}
