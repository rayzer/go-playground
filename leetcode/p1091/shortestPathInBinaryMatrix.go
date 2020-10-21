package main

/*
 * @lc app=leetcode id=1091 lang=golang
 *
 * [1091] Shortest Path in Binary Matrix
 */

//dp
//BFS
//A*

// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	if grid[0][0] == 1 && grid[c-1][c-1] == 1 {
		return -1
	}

}

// @lc code=end

func main() {

}
