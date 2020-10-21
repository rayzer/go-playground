package main

import (
	"fmt"
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start

//f(1)

func coinChangeR(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	var helper func(idx, amount int) int
	helper = func(idx, amount int) int {
		if amount == 0 {
			return 0
		}
		if idx < len(coins) && amount > 0 {
			maxVal := amount / coins[idx]
			minCost := math.MaxInt64
			for i := 0; i <= maxVal; i++ {
				if amount >= i*coins[idx] {
					result := helper(idx+1, amount-i*coins[idx])
					if result != -1 {
						minCost = lmin(minCost, result+i)
					}
				}
			}
			if minCost == math.MaxInt64 {
				return -1
			} else {
				return minCost
			}
		}
		return -1
	}
	return helper(0, amount)
}

func lmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	fmt.Println(coinChangeR([]int{1}, 2))
	fmt.Println(coinChangeR([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChangeR([]int{1, 2, 5, 10}, 27))
}
