package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 */

/*

base case：
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

i: 第 n 天
k： 还允许 k 次交易
0: 手上无股票
1：手上有股票
*/

// @lc code=start
func maxProfit(k int, prices []int) int {
	dp := make([][][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		for nk := 0; nk < k+1; nk++ {
			tmp := []int{0, 0}
			dp[i] = append(dp[i], tmp)
		}
	}

	fmt.Println(dp)
	for i := 0; i < len(dp); i++ {
		for kn := k; kn >= 1; kn-- {
			fmt.Println(i, kn)
			if i-1 == -1 {
				continue
			}
			dp[i][kn][0] = max(dp[i-1][kn][0], dp[i-1][kn-1][1]+prices[i])
			dp[i][kn][1] = max(dp[i-1][kn][1], dp[i-1][kn-1][0]-prices[i])
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))          //2
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3})) //7
}
