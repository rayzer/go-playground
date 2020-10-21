package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Friend Circles
 */

// @lc code=start
func findCircleNum2(M [][]int) int {
	visited := make([]int, len(M))

	var dfs func([][]int, []int, int)
	dfs = func(M [][]int, visited []int, person int) {
		for other := 0; other < len(M); other++ {
			if M[person][other] == 1 && visited[other] == 0 {
				visited[other] = 1
				dfs(M, visited, other)
			}
		}
	}

	count := 0
	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			dfs(M, visited, i)
			count++
		}
	}
	return count
}

// @lc code=end

func main() {
	input := [][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	}
	fmt.Println(findCircleNum2(input)) //2
}
