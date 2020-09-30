package main

/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
func totalNQueens(n int) int {
	var total int
	col := make([]int, n)
	diag := make([]int, n*2)
	diag2 := make([]int, n*2)
	places := make([]int, n)

	// recursive backtracking function
	var search func(y int)
	search = func(y int) {
		if y == n {
			total++
			return
		}

		for x := 0; x < n; x++ {
			if col[x] != 0 || diag[x+y] != 0 || diag2[x-y+n-1] != 0 {
				continue
			}
			col[x], diag[x+y], diag2[x-y+n-1], places[x] = 1, 1, 1, y
			search(y + 1)
			col[x], diag[x+y], diag2[x-y+n-1], places[x] = 0, 0, 0, 0
		}

	}
	search(0)
	return total
}

// @lc code=end
