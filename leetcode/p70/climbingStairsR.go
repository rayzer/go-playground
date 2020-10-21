package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

//1: 1
//2: 2
//f(3) = f(2) + f(1)
//f(n) = f(n-1) + f(n-2)

func climbStairsR(n int) int {
	memory := make(map[int]int)
	memory[1] = 1
	memory[2] = 2

	var helper func(n int) int
	helper = func(n int) int {
		if _, ok := memory[n]; ok {
			return memory[n]
		}
		result := helper(n-1) + helper(n-2)
		memory[n] = result
		return result
	}
	return helper(n)
}

// @lc code=end

func main() {
	fmt.Println(climbStairsR(10))
}
