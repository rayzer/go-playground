package main

import "fmt"

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

// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a := 1
	b := 2
	result := 0
	for i := 3; i <= n; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(climbStairs(10))
}
