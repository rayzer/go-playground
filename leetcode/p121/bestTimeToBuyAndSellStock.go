package main

import "fmt"

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i; j < len(prices); j++ {
			if prices[j] > prices[i] {
				profit := prices[j] - prices[i]
				if profit > max {
					max = profit
				}
			}
		}
	}
	return max
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
