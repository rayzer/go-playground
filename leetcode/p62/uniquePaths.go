package main

import "fmt"

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
//time limit exceed
//func uniquePaths(m int, n int) int {
//	if m <= 0 || n <= 0 {
//		return 0
//	}
//	if m == 1 || n == 1 {
//		return 1
//	}
//	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
//}

//beats 100 % of golang submissions
//func uniquePaths(m int, n int) int {
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, n)
//	}
//	for i := 0; i < n; i++ {
//		dp[0][i] = 1
//	}
//	for i := 0; i < m; i++ {
//		dp[i][0] = 1
//	}
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			dp[i][j] = dp[i-1][j] + dp[i][j-1]
//		}
//	}
//	return dp[m-1][n-1]
//}

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[n-1]
}

// @lc code=end

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 3))
}
