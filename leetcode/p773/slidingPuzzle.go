package main

import "fmt"

/*
 * @lc app=leetcode id=773 lang=golang
 *
 * [773] Sliding Puzzle
 */

//DFS
//BFS
//A *

// @lc code=start
func slidingPuzzle(board [][]int) int {
	moves := make(map[int][]int)
	moves[0] = []int{1, 3}
	moves[1] = []int{0, 2, 4}
	moves[2] = []int{1, 5}
	moves[3] = []int{0, 4}
	moves[4] = []int{1, 3, 5}
	moves[5] = []int{2, 4}

}

// @lc code=end

func main() {
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
}
