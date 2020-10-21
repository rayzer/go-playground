package main

import "fmt"

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
//0 ms, faster than 100.00% of Go

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var m, n int
	m = len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n = len(obstacleGrid[0])
	//if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
	//	return 0
	//}

	dp := make([]int, n)
	dp[0] = 1
	//for i := 0; i < n; i++ {
	//	if obstacleGrid[0][i] == 1 {
	//		for j := i; j < n; j++ {
	//			dp[j] = 0
	//		}
	//		break
	//	} else {
	//		dp[i] = 1
	//	}
	//}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				if j > 0 {
					dp[j] = dp[j] + dp[j-1]
				}
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
	fmt.Println(uniquePathsWithObstacles(input))
}
