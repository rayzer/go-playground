package main

import "fmt"

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

//分治：求每一行最小，相加
//暴力：累加求最小
//

// @lc code=start
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	fmt.Println(triangle)
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
func main() {
	input := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(input))
}
