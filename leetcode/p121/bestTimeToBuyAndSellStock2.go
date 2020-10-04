package main

import "fmt"

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
//4 ms, faster than 93.60%
func maxProfit(prices []int) int {
	max, min := 0, prices[0]
	for _, v := range prices {
		temp := v - min
		if temp > max {
			max = temp
		}
		if temp < 0 {
			temp = 0
		}
		//keep smallest passby
		if min > v {
			min = v
		}
	}
	return max
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
