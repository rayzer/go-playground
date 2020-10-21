package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change


1. 暴力递归
	每次递归分别尝试减去1，2，5，直到1
2. BFS
3. DP


*/

//DP: f(n) = min(f(n-k), k in {1,2,5}) + 1
// @lc code=start

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

//BFS
func coinChange2(coins []int, amount int) int {
	if len(coins) == 0 || amount <= 0 {
		return 0
	}
	coinsMap := make(map[int]bool)
	for _, coin := range coins {
		coinsMap[coin] = true
	}

	visited := make(map[int]bool)
	level := []int{amount}

	var count int
	var temp []int

	for len(level) > 0 {
		count += 1
		temp = []int{}
		for _, a := range level {
			if coinsMap[a] == true {
				return count
			}
			for _, c := range coins {
				if a-c > 0 && visited[a-c] == false {
					temp = append(temp, a-c)
					visited[a-c] = true
				}
			}
		}
		level = temp
	}
	return -1
}

func main() {
	fmt.Println(coinChange([]int{1}, 2))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	//fmt.Println(coinChange2([]int{186, 419, 83, 408}, 6249))
}
