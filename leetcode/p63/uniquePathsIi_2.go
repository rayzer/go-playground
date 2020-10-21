package main

import "fmt"

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	dp[0] = 1
	for _, row := range obstacleGrid {
		for i := 0; i < n; i++ {
			if row[i] == 1 {
				dp[i] = 0
			} else if i > 0 {
				dp[i] = dp[i] + dp[i-1]
			}
		}
	}
	return dp[n-1]
}

// @lc code=end
func main() {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	//input := [][]int{{0, 0}, {1, 0}}
	//input := [][]int{{0, 1, 0, 0, 0}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles2(input))
}
